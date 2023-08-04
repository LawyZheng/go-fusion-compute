package storage

import (
	"encoding/json"
	"strings"

	"github.com/lawyzheng/go-fusion-compute/client"
	"github.com/lawyzheng/go-fusion-compute/internal/common"
)

const (
	siteMask     = "<site_uri>"
	datastoreUrl = "<site_uri>/datastores"
)

type Interface interface {
	ListDataStore() ([]Datastore, error)
}

func NewManager(client client.FusionComputeClient, siteUri string) Interface {
	return &manager{client: client, siteUri: siteUri}
}

type manager struct {
	client  client.FusionComputeClient
	siteUri string
}

func (m *manager) ListDataStore() ([]Datastore, error) {
	var adapters []Datastore
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().Get(strings.Replace(datastoreUrl, siteMask, m.siteUri, -1))
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		var listAdapterResponse ListDataStoreResponse
		err := json.Unmarshal(resp.Body(), &listAdapterResponse)
		if err != nil {
			return nil, err
		}
		adapters = listAdapterResponse.Datastores
	} else {
		return nil, common.FormatHttpError(resp)
	}
	return adapters, nil
}
