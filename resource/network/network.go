package network

import (
	"context"
	"path"
	"strings"

	"github.com/lawyzheng/go-fusion-compute/client"
	"github.com/lawyzheng/go-fusion-compute/resource/vm"
)

const (
	siteMask     = "<site_uri>"
	dvSwitchUrl  = "<site_uri>/dvswitchs"
	vmScopeUrl   = "<site_uri>/vms?scope=<resource_urn>"
	portGroupUrl = "<site_uri>/portgroups"
)

var _ Manager = (*manager)(nil)

type Manager interface {
	ListDVSwitch(ctx context.Context, siteUri string) ([]DVSwitch, error)
	ListPortGroupBySwitch(ctx context.Context, dvSwitchIdUri string) ([]PortGroup, error)
	ListPortGroupInUseIp(ctx context.Context, siteUri string, portGroupUrn string) ([]string, error)
	ListPortGroup(ctx context.Context, siteUri string) ([]PortGroup, error)
}

func NewManager(client client.FusionComputeClient) Manager {
	return &manager{client: client}
}

type manager struct {
	client client.FusionComputeClient
}

func (m *manager) ListPortGroup(ctx context.Context, siteUri string) ([]PortGroup, error) {
	uri := strings.Replace(portGroupUrl, siteMask, siteUri, -1)
	listPortGroupResponse := new(ListPortGroupResponse)
	if err := client.Get(ctx, m.client, uri, listPortGroupResponse); err != nil {
		return nil, err
	}
	return listPortGroupResponse.PortGroups, nil
}

func (m *manager) ListPortGroupBySwitch(ctx context.Context, dvSwitchIdUri string) ([]PortGroup, error) {
	uri := path.Join(dvSwitchIdUri, "portgroups")
	listPortGroupResponse := new(ListPortGroupResponse)
	if err := client.Get(ctx, m.client, uri, listPortGroupResponse); err != nil {
		return nil, err
	}
	return listPortGroupResponse.PortGroups, nil
}

func (m *manager) ListDVSwitch(ctx context.Context, siteUri string) ([]DVSwitch, error) {
	uri := strings.Replace(dvSwitchUrl, siteMask, siteUri, -1)
	listDVSwitchResponse := new(ListDVSwitchResponse)
	if err := client.Get(ctx, m.client, uri, listDVSwitchResponse); err != nil {
		return nil, err
	}
	return listDVSwitchResponse.DVSwitchs, nil
}

func (m *manager) ListPortGroupInUseIp(ctx context.Context, siteUri string, portGroupUrn string) ([]string, error) {
	uri := strings.Replace(strings.Replace(vmScopeUrl, siteMask, siteUri, -1), "<resource_urn>", portGroupUrn, -1)
	listVmResponse := new(vm.ListVmResponse)
	if err := client.Get(ctx, m.client, uri, listVmResponse); err != nil {
		return nil, err
	}

	var results []string
	for _, v := range listVmResponse.Vms {
		for _, nic := range v.VmConfig.Nics {
			if nic.IP != "0.0.0.0" {
				results = append(results, nic.IP)
			}
		}
	}
	return results, nil
}
