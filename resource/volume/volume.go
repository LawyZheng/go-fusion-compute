package volume

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"

	"github.com/lawyzheng/go-fusion-compute/client"
)

var _ Manager = (*rbdManager)(nil)

type Manager interface {
	GetVolume(ctx context.Context, volumeUri string) (*Volume, error)
	ExportVolume(ctx context.Context, volumeUrl string, opt *ExportOption) error
	ImportVolume(ctx context.Context, volumeUrl string, opt *ImportOption) error
	MergeVolume(ctx context.Context, opt *MergeOption) error
}

func newManager(client client.FusionComputeClient) *manager {
	return &manager{client: client}
}

type manager struct {
	client client.FusionComputeClient
}

func (m *manager) GetVolume(ctx context.Context, volumeUri string) (*Volume, error) {
	vol := new(Volume)
	if err := client.Get(ctx, m.client, volumeUri, vol); err != nil {
		return nil, err
	}
	return vol, nil
}

func NewRBDManager(client client.FusionComputeClient) Manager {
	return &rbdManager{
		manager: newManager(client),
	}
}

type rbdManager struct {
	*manager
}

func (m *rbdManager) ExportVolume(ctx context.Context, volumeUrl string, opt *ExportOption) error {
	if err := opt.Validate(); err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, "rbd", "export-diff")
	stderr := new(bytes.Buffer)
	cmd.Stderr = stderr

	cmd.Args = append(cmd.Args, fmt.Sprintf("%s@%s", volumeUrl, *opt.CurrentSnapshot))
	if opt.FromSnapshot != nil {
		cmd.Args = append(cmd.Args, "--from-snap", *opt.FromSnapshot)
	}
	if opt.FilePath != nil {
		cmd.Args = append(cmd.Args, *opt.FilePath)
	} else {
		cmd.Args = append(cmd.Args, "-", "--no-progress") // output to stdout
		cmd.Stdout = opt.Writer
	}

	if err := cmd.Run(); err != nil {
		if cmd.ProcessState == nil {
			return err
		}
		return fmt.Errorf("%s; error: %s", cmd.ProcessState.String(), stderr)
	}

	if code := cmd.ProcessState.ExitCode(); code != 0 {
		return fmt.Errorf("%s; error: %s", cmd.ProcessState.String(), stderr)
	}

	return nil
}

func (m *rbdManager) ImportVolume(ctx context.Context, volumeUrl string, opt *ImportOption) error {
	if err := opt.Validate(); err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, "rbd", "import-diff")
	if opt.FilePath != nil {
		cmd.Args = append(cmd.Args, *opt.FilePath, volumeUrl)
	} else {
		cmd.Args = append(cmd.Args, "-", volumeUrl)
		stdin, err := cmd.StdinPipe()
		if err != nil {
			return err
		}
		go func() {
			defer stdin.Close()
			if _, err := io.Copy(stdin, opt.Reader); err != nil {
				io.WriteString(cmd.Stderr, err.Error())
			}
		}()
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		if cmd.ProcessState == nil {
			return err
		}
		return fmt.Errorf("%s; error:%s", cmd.ProcessState.String(), string(output))
	}
	if cmd.ProcessState.ExitCode() != 0 {
		return fmt.Errorf("%s; error:%s", cmd.ProcessState.String(), string(output))
	}
	return nil
}

func (m *rbdManager) MergeVolume(ctx context.Context, opt *MergeOption) error {
	if err := opt.Validate(); err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, "rbd", "merge-diff")
	stderr := new(bytes.Buffer)
	cmd.Stderr = stderr

	if opt.SrcPath != nil {
		cmd.Args = append(cmd.Args, *opt.SrcPath)
	} else {
		cmd.Args = append(cmd.Args, "-")

		stdin, err := cmd.StdinPipe()
		if err != nil {
			return err
		}

		go func() {
			defer stdin.Close()
			if _, err := io.Copy(stdin, opt.SrcReader); err != nil {
				io.WriteString(cmd.Stderr, err.Error())
			}
		}()

	}

	cmd.Args = append(cmd.Args, *opt.DiffPath)

	if opt.DstPath != nil {
		cmd.Args = append(cmd.Args, *opt.DstPath)
	} else {
		cmd.Args = append(cmd.Args, "-", "--no-progress")
		cmd.Stdout = opt.DstWriter
	}

	if err := cmd.Run(); err != nil {
		if cmd.ProcessState == nil {
			return err
		}
		return fmt.Errorf("%s; error: %s", cmd.ProcessState.String(), stderr)
	}

	if code := cmd.ProcessState.ExitCode(); code != 0 {
		return fmt.Errorf("%s; error: %s", cmd.ProcessState.String(), stderr)
	}

	return nil
}
