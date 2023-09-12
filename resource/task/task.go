package task

import (
	"context"
	"time"

	"github.com/lawyzheng/go-fusion-compute/client"
)

var _ Manager = (*manager)(nil)

type Manager interface {
	Get(ctx context.Context, taskUri string) (*Task, error)
	Wait(ctx context.Context, taskUri string, interval time.Duration) <-chan TaskResult
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

func (m *manager) Wait(ctx context.Context, taskUri string, interval time.Duration) <-chan TaskResult {
	ch := make(chan TaskResult, 1)
	go func() {
		defer close(ch)

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				task, err := m.Get(ctx, taskUri)
				if err != nil {
					ch <- TaskResult{
						Err: err,
					}
					return
				}

				ch <- TaskResult{
					Task: *task,
				}

				if task.IsDone() {
					return
				}

			case <-ctx.Done():
				ch <- TaskResult{
					Err: ctx.Err(),
				}
				return
			}
		}
	}()
	return ch
}
