package cluster

import (
	"encoding/json"
	"strings"

	"github.com/lawyzheng/go-fusion-compute/client"
	"github.com/lawyzheng/go-fusion-compute/internal/common"
	fcErr "github.com/lawyzheng/go-fusion-compute/pkg/error"
)

const (
	siteMask   = "<site_uri>"
	clusterUrl = "<site_uri>/clusters"
)

type Manager interface {
	ListCluster() ([]Cluster, error)
}

func NewManager(client client.FusionComputeClient, siteUri string) Manager {
	return &manager{client: client, siteUri: siteUri}
}

type manager struct {
	client  client.FusionComputeClient
	siteUri string
}

func (m *manager) ListCluster() ([]Cluster, error) {
	var clusters []Cluster
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().Get(strings.Replace(clusterUrl, siteMask, m.siteUri, -1))
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		var listClusterResponse ListClusterResponse
		err := json.Unmarshal(resp.Body(), &listClusterResponse)
		if err != nil {
			return nil, err
		}
		clusters = listClusterResponse.Clusters
	} else {
		e := new(fcErr.Basic)
		return nil, common.FormatHttpError(resp, e)
	}
	return clusters, nil

}
