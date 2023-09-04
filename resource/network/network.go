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

type Manager interface {
	ListDVSwitch(ctx context.Context) ([]DVSwitch, error)
	ListPortGroupBySwitch(ctx context.Context, dvSwitchIdUri string) ([]PortGroup, error)
	ListPortGroupInUseIp(ctx context.Context, portGroupUrn string) ([]string, error)
	ListPortGroup(ctx context.Context) ([]PortGroup, error)
}

func NewManager(client client.FusionComputeClient, siteUri string) Manager {
	return &manager{client: client, siteUri: siteUri}
}

type manager struct {
	client  client.FusionComputeClient
	siteUri string
}

func (m *manager) ListPortGroup(ctx context.Context) ([]PortGroup, error) {
	uri := strings.Replace(portGroupUrl, siteMask, m.siteUri, -1)
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

func (m *manager) ListDVSwitch(ctx context.Context) ([]DVSwitch, error) {
	uri := strings.Replace(dvSwitchUrl, siteMask, m.siteUri, -1)
	listDVSwitchResponse := new(ListDVSwitchResponse)
	if err := client.Get(ctx, m.client, uri, listDVSwitchResponse); err != nil {
		return nil, err
	}
	return listDVSwitchResponse.DVSwitchs, nil
}

func (m *manager) ListPortGroupInUseIp(ctx context.Context, portGroupUrn string) ([]string, error) {
	uri := strings.Replace(strings.Replace(vmScopeUrl, siteMask, m.siteUri, -1), "<resource_urn>", portGroupUrn, -1)
	listVmResponse := new(vm.ListVmResponse)
	if err := client.Get(ctx, m.client, uri, listVmResponse); err != nil {
		return nil, err
	}

	var results []string
	for _, v := range listVmResponse.Vms {
		for _, nic := range v.VmConfig.Nics {
			if nic.Ip != "0.0.0.0" {
				results = append(results, nic.Ip)
			}
		}
	}
	return results, nil
}
