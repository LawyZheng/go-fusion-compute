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

type Manager interface {
	ListDataStore(ctx context.Context) ([]Datastore, error)
}

func NewManager(client client.FusionComputeClient, siteUri string) Manager {
	return &manager{client: client, siteUri: siteUri}
}

type manager struct {
	client  client.FusionComputeClient
	siteUri string
}

func (m *manager) ListDataStore(ctx context.Context) ([]Datastore, error) {
	listAdapterResponse := new(ListDataStoreResponse)
	uri := strings.Replace(datastoreUrl, siteMask, m.siteUri, -1)
	if err := client.Get(ctx, m.client, uri, listAdapterResponse); err != nil {
		return nil, err
	}

	return listAdapterResponse.Datastores, nil
}
