package volume

import (
	"context"

	"github.com/lawyzheng/go-fusion-compute/client"
)

var _ Manager = (*manager)(nil)

type Manager interface {
	GetVolume(ctx context.Context, volumeUri string) (*Volume, error)
}

func NewManager(client client.FusionComputeClient) Manager {
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
