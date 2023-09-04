package site

import (
	"context"

	"github.com/lawyzheng/go-fusion-compute/client"
)

const (
	siteUrl = "/service/sites"
)

type Manager interface {
	ListSite(ctx context.Context) ([]Site, error)
	GetSite(ctx context.Context, siteUri string) (*Site, error)
}

func NewManager(client client.FusionComputeClient) Manager {
	return &manager{client: client}
}

type manager struct {
	client client.FusionComputeClient
}

func (m *manager) GetSite(ctx context.Context, siteUri string) (*Site, error) {
	site := new(Site)
	if err := client.Get(ctx, m.client, siteUri, site); err != nil {
		return nil, err
	}
	return site, nil
}

func (m *manager) ListSite(ctx context.Context) ([]Site, error) {
	listSiteResponse := new(ListSiteResponse)
	if err := client.Get(ctx, m.client, siteUrl, listSiteResponse); err != nil {
		return nil, err
	}
	return listSiteResponse.Sites, nil
}
