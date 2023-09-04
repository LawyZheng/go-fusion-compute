package snapshot

import (
	"context"
	"strings"

	"github.com/lawyzheng/go-fusion-compute/client"
)

const (
	vmMask      = "<vm_uri>"
	snapshotUrl = "<vm_uri>/snapshots"
)

type Manager interface {
	CreateSnapshot(ctx context.Context, req *CreateSnapshotReq) (*Task, error)
	DeleteSnapshot(ctx context.Context, snapshotUri string) (*Task, error)
	GetSnapshotDetail(ctx context.Context, snapshotUri string) (*SnapshotDetail, error)
	GetCurrentSnapshot(ctx context.Context) (*SnapshotBrief, error)
	ListSnapshots(ctx context.Context) (*ListSnapshotsResponse, error)
}

func NewManager(client client.FusionComputeClient, vmUri string) Manager {
	return &manager{vmUri: vmUri, client: client}
}

type manager struct {
	vmUri  string
	client client.FusionComputeClient
}

func (m *manager) CreateSnapshot(ctx context.Context, req *CreateSnapshotReq) (*Task, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	uri := strings.Replace(snapshotUrl, vmMask, m.vmUri, -1)

	task := new(Task)
	if err := client.Post(ctx, m.client, uri, req, task); err != nil {
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

func (m *manager) GetCurrentSnapshot(ctx context.Context) (*SnapshotBrief, error) {
	snapshots, err := m.ListSnapshots(ctx)
	if err != nil {
		return nil, err
	}

	return &snapshots.CurrentSnapshot, nil
}

func (m *manager) ListSnapshots(ctx context.Context) (*ListSnapshotsResponse, error) {
	uri := strings.Replace(snapshotUrl, vmMask, m.vmUri, -1)
	data := new(ListSnapshotsResponse)
	if err := client.Get(ctx, m.client, uri, data); err != nil {
		return nil, err
	}

	if len(data.RootSnapshots) == 0 {
		return nil, ErrNoSnapshot
	}

	return data, nil
}
