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

var _ Manager = (*manager)(nil)

type Manager interface {
	ListCluster(ctx context.Context, siteUri string) ([]Cluster, error)
}

func NewManager(client client.FusionComputeClient) Manager {
	return &manager{client: client}
}

type manager struct {
	client client.FusionComputeClient
}

func (m *manager) ListCluster(ctx context.Context, siteUri string) ([]Cluster, error) {
	uri := strings.Replace(clusterUrl, siteMask, siteUri, -1)
	listClusterResponse := new(ListClusterResponse)
	if err := client.Get(ctx, m.client, uri, listClusterResponse); err != nil {
		return nil, err
	}
	return listClusterResponse.Clusters, nil
}
