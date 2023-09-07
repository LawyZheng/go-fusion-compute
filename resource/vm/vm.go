package vm

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/lawyzheng/go-fusion-compute/client"
)

const (
	siteMask = "<site_uri>"
	vmUrl    = "<site_uri>/vms"
)

var _ Manager = (*manager)(nil)

type Manager interface {
	ListVm(ctx context.Context, siteUri string, isTemplate bool) ([]Vm, error)
	GetVM(ctx context.Context, vmUri string) (*Vm, error)
	CloneVm(ctx context.Context, templateUri string, request CloneVmRequest) (*CloneVmResponse, error)
	DeleteVm(ctx context.Context, vmUri string) (*DeleteVmResponse, error)
	UploadImage(ctx context.Context, vmUri string, request ImportTemplateRequest) (*ImportTemplateResponse, error)
}

func NewManager(client client.FusionComputeClient) Manager {
	return &manager{client: client}
}

type manager struct {
	client client.FusionComputeClient
}

func (m *manager) CloneVm(ctx context.Context, templateUri string, request CloneVmRequest) (*CloneVmResponse, error) {
	for n := range request.VmCustomization.NicSpecification {
		if !strings.Contains(request.VmCustomization.NicSpecification[n].Netmask, ".") {
			b, err := strconv.Atoi(request.VmCustomization.NicSpecification[0].Netmask)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("can not parse netmask: %s", err.Error()))
			}
			mask, err := parseMask(b)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("can not parse netmask: %s", err.Error()))
			}
			request.VmCustomization.NicSpecification[n].Netmask = mask
		}
	}

	vm, err := m.GetVM(ctx, templateUri)
	if err != nil {
		return nil, err
	}
	disks := vm.VmConfig.Disks
	if len(disks) > 0 {
		if request.Config.Disks[0].QuantityGB < vm.VmConfig.Disks[0].QuantityGB {
			request.Config.Disks[0].QuantityGB = vm.VmConfig.Disks[0].QuantityGB
		}
	}
	cloneVmResponse := new(CloneVmResponse)
	uri := path.Join(templateUri, "action", "clone")
	if err := client.Post(ctx, m.client, uri, &request, cloneVmResponse); err != nil {
		return nil, err
	}
	return cloneVmResponse, nil
}

func (m *manager) ListVm(ctx context.Context, siteUri string, isTemplate bool) ([]Vm, error) {
	u := new(url.URL)
	u.Path = strings.Replace(vmUrl, siteMask, siteUri, -1)
	if isTemplate {
		v := url.Values{}
		v.Set("isTemplate", "true")
		u.RawQuery = v.Encode()
	}

	listVmResponse := new(ListVmResponse)
	if err := client.Get(ctx, m.client, u.String(), listVmResponse); err != nil {
		return nil, err
	}
	return listVmResponse.Vms, nil
}

func (m *manager) DeleteVm(ctx context.Context, vmUri string) (*DeleteVmResponse, error) {
	deleteVmResponse := new(DeleteVmResponse)
	if err := client.Delete(ctx, m.client, vmUri, deleteVmResponse); err != nil {
		return nil, err
	}
	return deleteVmResponse, nil
}

func (m *manager) GetVM(ctx context.Context, vmUri string) (*Vm, error) {
	vm := new(Vm)
	if err := client.Get(ctx, m.client, vmUri, vm); err != nil {
		return nil, err
	}
	return vm, nil
}

func (m *manager) UploadImage(ctx context.Context, vmUri string, request ImportTemplateRequest) (*ImportTemplateResponse, error) {
	res := new(ImportTemplateResponse)
	uri := path.Join(vmUri, "action", "import")
	if err := client.Post(ctx, m.client, uri, &request, res); err != nil {
		return nil, err
	}
	return res, nil
}

func parseMask(num int) (mask string, err error) {
	var buff bytes.Buffer
	for i := 0; i < int(num); i++ {
		buff.WriteString("1")
	}
	for i := num; i < 32; i++ {
		buff.WriteString("0")
	}
	masker := buff.String()
	a, _ := strconv.ParseUint(masker[:8], 2, 64)
	b, _ := strconv.ParseUint(masker[8:16], 2, 64)
	c, _ := strconv.ParseUint(masker[16:24], 2, 64)
	d, _ := strconv.ParseUint(masker[24:32], 2, 64)
	resultMask := fmt.Sprintf("%v.%v.%v.%v", a, b, c, d)
	return resultMask, nil
}
