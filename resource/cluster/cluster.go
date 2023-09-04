package cluster

import (
	"context"
	"strings"

	"github.com/lawyzheng/go-fusion-compute/client"
)

const (
	siteMask   = "<site_uri>"
	clusterUrl = "<site_uri>/clusters"
)

type Manager interface {
	ListCluster(ctx context.Context) ([]Cluster, error)
}

func NewManager(client client.FusionComputeClient, siteUri string) Manager {
	return &manager{client: client, siteUri: siteUri}
}

type manager struct {
	client  client.FusionComputeClient
	siteUri string
}

func (m *manager) ListCluster(ctx context.Context) ([]Cluster, error) {
	uri := strings.Replace(clusterUrl, siteMask, m.siteUri, -1)
	listClusterResponse := new(ListClusterResponse)
	if err := client.Get(ctx, m.client, uri, listClusterResponse); err != nil {
		return nil, err
	}
	return listClusterResponse.Clusters, nil
}
