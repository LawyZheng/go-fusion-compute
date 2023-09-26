package rbd

import (
	"context"
)

var (
	rbd = New()
)

func DeleteSnapshot(ctx context.Context, volumeUrl string, snapId string) error {
	return rbd.DeleteSnapshot(ctx, volumeUrl, snapId)
}

func ExportVolume(ctx context.Context, volumeUrl string, opt *ExportOption) error {
	return rbd.ExportVolume(ctx, volumeUrl, opt)
}

func ImportVolume(ctx context.Context, volumeUrl string, opt *ImportOption) error {
	return rbd.ImportVolume(ctx, volumeUrl, opt)
}

func MergeVolume(ctx context.Context, opt *MergeOption) error {
	return rbd.MergeVolume(ctx, opt)
}
