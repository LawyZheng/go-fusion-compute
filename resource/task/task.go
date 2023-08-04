package task

import (
	"encoding/json"
	"path"

	"github.com/lawyzheng/go-fusion-compute/client"
	"github.com/lawyzheng/go-fusion-compute/internal/common"
)

type Manager interface {
	Get(taskUri string) (*Task, error)
}

func NewManager(client client.FusionComputeClient, siteUri string) Manager {
	return &manager{client: client, siteUri: siteUri}
}

type manager struct {
	client  client.FusionComputeClient
	siteUri string
}

func (m *manager) Get(taskUri string) (*Task, error) {
	var task Task
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().Get(path.Join(taskUri))
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		err := json.Unmarshal(resp.Body(), &task)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, common.FormatHttpError(resp)
	}
	return &task, nil
}
