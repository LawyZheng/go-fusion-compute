package rbd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"
)

func DeleteSnapshot(ctx context.Context, volumeUrl string, snapId string) error {
	snap := fmt.Sprintf("%s@%s", volumeUrl, snapId)
	cmd := exec.CommandContext(ctx, "rbd", "snap", "rm", snap)
	return combinedOutput(cmd)
}

func ExportVolume(ctx context.Context, volumeUrl string, opt *ExportOption) error {
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

	return runWithStderr(cmd, stderr)
}

func ImportVolume(ctx context.Context, volumeUrl string, opt *ImportOption) error {
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

	return combinedOutput(cmd)
}

func MergeVolume(ctx context.Context, opt *MergeOption) error {
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

	return runWithStderr(cmd, stderr)
}

func runWithStderr(cmd *exec.Cmd, stderr fmt.Stringer) error {
	if err := cmd.Run(); err != nil {
		if cmd.ProcessState == nil {
			return err
		}
		return fmt.Errorf("%s; error: %s", cmd.ProcessState.String(), stderr.String())
	}

	if code := cmd.ProcessState.ExitCode(); code != 0 {
		return fmt.Errorf("%s; error: %s", cmd.ProcessState.String(), stderr.String())
	}

	return nil
}

func combinedOutput(cmd *exec.Cmd) error {
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