package task

import (
	"context"

	"github.com/lawyzheng/go-fusion-compute/client"
)

var _ Manager = (*manager)(nil)

type Manager interface {
	Get(ctx context.Context, taskUri string) (*Task, error)
}

func NewManager(client client.FusionComputeClient) Manager {
	return &manager{client: client}
}

type manager struct {
	client client.FusionComputeClient
}

func (m *manager) Get(ctx context.Context, taskUri string) (*Task, error) {
	task := new(Task)
	if err := client.Get(ctx, m.client, taskUri, task); err != nil {
		return nil, err
	}
	return task, nil
}
