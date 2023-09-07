package storage

import (
	"context"
	"strings"

	"github.com/lawyzheng/go-fusion-compute/client"
)

const (
	siteMask     = "<site_uri>"
	datastoreUrl = "<site_uri>/datastores"
)

var _ Manager = (*manager)(nil)

type Manager interface {
	ListDataStore(ctx context.Context, siteUri string) ([]Datastore, error)
}

func NewManager(client client.FusionComputeClient) Manager {
	return &manager{client: client}
}

type manager struct {
	client client.FusionComputeClient
}

func (m *manager) ListDataStore(ctx context.Context, siteUri string) ([]Datastore, error) {
	listAdapterResponse := new(ListDataStoreResponse)
	uri := strings.Replace(datastoreUrl, siteMask, siteUri, -1)
	if err := client.Get(ctx, m.client, uri, listAdapterResponse); err != nil {
		return nil, err
	}

	return listAdapterResponse.Datastores, nil
}
