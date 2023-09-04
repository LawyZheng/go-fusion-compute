package host

import (
	"context"
	"strings"

	"github.com/lawyzheng/go-fusion-compute/client"
)

const (
	siteMask = "<site_uri>"
	hostUrl  = "<site_uri>/hosts"
)

type Manager interface {
	ListHosts(ctx context.Context) ([]Host, error)
}

func NewManager(siteUri string, client client.FusionComputeClient) Manager {
	return &manager{siteUri: siteUri, client: client}
}

type manager struct {
	siteUri string
	client  client.FusionComputeClient
}

func (m *manager) ListHosts(ctx context.Context) ([]Host, error) {
	uri := strings.Replace(hostUrl, siteMask, m.siteUri, -1)

	list := new(listHostsResponse)
	if err := client.Get(ctx, m.client, uri, list); err != nil {
		return nil, err
	}

	return list.Hosts, nil
}
