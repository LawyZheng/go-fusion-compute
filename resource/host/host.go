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

var _ Manager = (*manager)(nil)

type Manager interface {
	ListHosts(ctx context.Context, siteUri string) ([]Host, error)
}

func NewManager(client client.FusionComputeClient) Manager {
	return &manager{client: client}
}

type manager struct {
	client client.FusionComputeClient
}

func (m *manager) ListHosts(ctx context.Context, siteUri string) ([]Host, error) {
	uri := strings.Replace(hostUrl, siteMask, siteUri, -1)

	list := new(ListHostsResponse)
	if err := client.Get(ctx, m.client, uri, list); err != nil {
		return nil, err
	}

	return list.Hosts, nil
}
