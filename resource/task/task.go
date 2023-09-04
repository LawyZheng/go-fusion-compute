package task

import (
	"context"

	"github.com/lawyzheng/go-fusion-compute/client"
)

type Manager interface {
	Get(ctx context.Context, taskUri string) (*Task, error)
}

func NewManager(client client.FusionComputeClient, siteUri string) Manager {
	return &manager{client: client, siteUri: siteUri}
}

type manager struct {
	client  client.FusionComputeClient
	siteUri string
}

func (m *manager) Get(ctx context.Context, taskUri string) (*Task, error) {
	task := new(Task)
	if err := client.Get(ctx, m.client, taskUri, task); err != nil {
		return nil, err
	}
	return task, nil
}
