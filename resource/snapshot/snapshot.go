package snapshot

import (
	"context"

	"github.com/lawyzheng/go-fusion-compute/client"
)

const (
	vmMask      = "<vm_uri>"
	snapshotUrl = "<vm_uri>/snapshots"
)

type Manager interface {
	CreateSnapshot(ctx context.Context, vmUri string, req *CreateSnapshotReq) (*Task, error)
	DeleteSnapshot(ctx context.Context, snapshotUri string) (*Task, error)
	GetSnapshotDetail(ctx context.Context, snapshotUri string) (*SnapshotDetail, error)
	GetCurrentSnapshot(ctx context.Context, vmUri string) (*SnapshotBrief, error)
	ListSnapshots(ctx context.Context, vmUri string) (*ListSnapshotsResponse, error)
}

func NewManager(client client.FusionComputeClient) Manager {
	return &manager{client: client}
}

type manager struct {
	client client.FusionComputeClient
}

func (m *manager) CreateSnapshot(ctx context.Context, vmUri string, req *CreateSnapshotReq) (*Task, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	task := new(Task)
	if err := client.Post(ctx, m.client, vmUri, req, task); err != nil {
		return nil, err
	}

	return task, nil
}

func (m *manager) DeleteSnapshot(ctx context.Context, snapshotUri string) (*Task, error) {
	task := new(Task)
	if err := client.Delete(ctx, m.client, snapshotUri, task); err != nil {
		return nil, err
	}
	return task, nil
}

func (m *manager) GetSnapshotDetail(ctx context.Context, snapshotUri string) (*SnapshotDetail, error) {
	snapshot := new(SnapshotDetail)
	if err := client.Get(ctx, m.client, snapshotUri, snapshot); err != nil {
		return nil, err
	}
	return snapshot, nil
}

func (m *manager) GetCurrentSnapshot(ctx context.Context, vmUri string) (*SnapshotBrief, error) {
	snapshots, err := m.ListSnapshots(ctx, vmUri)
	if err != nil {
		return nil, err
	}

	return &snapshots.CurrentSnapshot, nil
}

func (m *manager) ListSnapshots(ctx context.Context, vmUri string) (*ListSnapshotsResponse, error) {
	data := new(ListSnapshotsResponse)
	if err := client.Get(ctx, m.client, vmUri, data); err != nil {
		return nil, err
	}

	if len(data.RootSnapshots) == 0 {
		return nil, ErrNoSnapshot
	}

	return data, nil
}
