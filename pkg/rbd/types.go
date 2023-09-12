package rbd

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var (
	ErrEmptyCurrentSnapshot = errors.New("current snapshot can't be empty")
	ErrRelativeFilePath     = errors.New("filepath is relative")
	ErrNoFilePath           = errors.New("export file path should be specified")
	ErrExistedFilePath      = errors.New("export file path shouldn't be an existing file")
	ErrNoDiffPath           = errors.New("diff path can't be empty")
)

func validateWriter(file *string, writer io.Writer) error {
	if file == nil || strings.TrimSpace(*file) == "" {
		if writer == nil || reflect.ValueOf(writer).IsNil() {
			return ErrNoFilePath
		}
	} else {
		p := strings.TrimSpace(*file)
		if !filepath.IsAbs(p) {
			return ErrRelativeFilePath
		}
		// FilePath shouldn't be existed
		_, err := os.Stat(p)
		if err == nil {
			return ErrExistedFilePath
		}

		if !os.IsNotExist(err) {
			return err
		}
	}

	return nil
}

func validateReader(file *string, reader io.Reader) error {
	if file == nil || strings.TrimSpace(*file) == "" {
		if reader == nil || reflect.ValueOf(reader).IsNil() {
			return ErrNoFilePath
		}
	} else {
		p := strings.TrimSpace(*file)
		if !filepath.IsAbs(p) {
			return ErrRelativeFilePath
		}
	}

	return nil
}

func NewExportOption(currentSnapshot string) *ExportOption {
	return &ExportOption{CurrentSnapshot: &currentSnapshot}
}

type ExportOption struct {
	CurrentSnapshot *string
	FromSnapshot    *string
	Writer          io.Writer
	// "FilePath" comes first, then "Writer"
	// "FilePath" shouldn't be existed and  should be an absolute path
	FilePath *string
}

func (op *ExportOption) SetWriter(writer io.Writer) *ExportOption {
	op.Writer = writer
	return op
}

// SetFilePath setting export filepath, the path can't be an existing object
func (op *ExportOption) SetFilePath(file string) *ExportOption {
	op.FilePath = &file
	return op
}

func (op *ExportOption) SetFromSnapshot(snapshotId string) *ExportOption {
	op.FromSnapshot = &snapshotId
	return op
}

func (op *ExportOption) Validate() error {
	if op.CurrentSnapshot == nil || strings.TrimSpace(*op.CurrentSnapshot) == "" {
		return ErrEmptyCurrentSnapshot
	}

	return validateWriter(op.FilePath, op.Writer)
}

func NewImportOption() *ImportOption {
	return &ImportOption{}
}

type ImportOption struct {
	Reader io.Reader
	// "FilePath" comes first, then "Reader"
	// "FilePath" should be an absolute path
	FilePath *string
}

func (op *ImportOption) SetReader(reader io.Reader) *ImportOption {
	op.Reader = reader
	return op
}

// SetFilePath setting export filepath, the path need to be an existing object
func (op *ImportOption) SetFilePath(file string) *ImportOption {
	op.FilePath = &file
	return op
}

func (op *ImportOption) Validate() error {
	return validateReader(op.FilePath, op.Reader)
}

func NewMergeOption(diff string) *MergeOption {
	return &MergeOption{DiffPath: &diff}
}

type MergeOption struct {
	DiffPath *string
	// "SrcPath" comes first, then "SrcReader"
	// "SrcPath" should be an absolute path
	SrcPath   *string
	SrcReader io.Reader
	// "DstPath" comes first, then "DstWriter"
	// "DstPath" shouldn't be existed and should be an absolute path
	DstPath   *string
	DstWriter io.Writer
}

func (op *MergeOption) SetDstWriter(writer io.Writer) *MergeOption {
	op.DstWriter = writer
	return op
}

// SetDstPath setting export filepath, the path can't be an existing object
func (op *MergeOption) SetDstPath(file string) *MergeOption {
	op.DstPath = &file
	return op
}

func (op *MergeOption) SetSrcReader(reader io.Reader) *MergeOption {
	op.SrcReader = reader
	return op
}

// SetSrcPath setting export filepath, the path need to be an existing object
func (op *MergeOption) SetSrcPath(file string) *MergeOption {
	op.SrcPath = &file
	return op
}

func (op *MergeOption) Validate() error {
	if op.DiffPath == nil {
		return ErrNoDiffPath
	}
	if !filepath.IsAbs(*op.DiffPath) {
		return ErrRelativeFilePath
	}

	if err := validateReader(op.SrcPath, op.SrcReader); err != nil {
		return err
	}

	return validateWriter(op.DstPath, op.DstWriter)
}
