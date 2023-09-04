package snapshot

import (
	"errors"
	"strings"
)

var (
	ErrNoSnapshot = errors.New("no snapshot found")
)

type SnapshotBrief struct {
	Urn            string          `json:"urn"`
	Uri            string          `json:"uri"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	CreateTime     string          `json:"createTime"`
	Status         string          `json:"status"`
	Type           string          `json:"type"`
	ChildSnapshots []SnapshotBrief `json:"childSnapshots"`
}

type ListSnapshotsResponse struct {
	CurrentSnapshot SnapshotBrief   `json:"currentSnapshot"`
	RootSnapshots   []SnapshotBrief `json:"rootSnapshots"`
}

type SnapshotDetail struct {
	Urn          string `json:"urn"`
	Uri          string `json:"uri"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CreateTime   string `json:"createTime"`
	Status       string `json:"status"`
	Type         string `json:"type"`
	VolSnapshots []struct {
		VolumeUrn      string `json:"volumeUrn"`
		VolumeUri      string `json:"volumeUri"`
		SnapUUID       string `json:"snapUuid"`
		StorageType    string `json:"storageType"`
		DatastoreUrn   string `json:"datastoreUrn"`
		SnapNameOneDev string `json:"snapNameOnDev"`
		ChgID          string `json:"chgID"`
	} `json:"volsnapshots"`
	SnapProvisionSize       int64 `json:"snapProvisionSize"`
	CoreNum                 int64 `json:"coreNum"`
	MemorySize              int64 `json:"memorySize"`
	VolumeSizeSum           int64 `json:"volumeSizeSum"`
	IncludingMemorySnapshot bool  `json:"includingMemorySnapshot"`
}

func NewCreateSnapshotReq(name string) *CreateSnapshotReq {
	return &CreateSnapshotReq{Name: &name}
}

type CreateSnapshotReq struct {
	NeedMemoryShot *bool   `json:"needMemoryShot,omitempty"`
	IsConsistent   *bool   `json:"isConsistent,omitempty"` // reserved field
	Description    *string `json:"description,omitempty"`
	Name           *string `json:"name"`
}

func (req *CreateSnapshotReq) SetDescription(desc string) *CreateSnapshotReq {
	req.Description = &desc
	return req
}

func (req *CreateSnapshotReq) SetNeedMemoryShot(ok bool) *CreateSnapshotReq {
	req.NeedMemoryShot = &ok
	return req
}

func (req *CreateSnapshotReq) Validate() error {
	if req.Name == nil || strings.TrimSpace(*req.Name) == "" {
		return errors.New("name can't be empty")
	}
	return nil
}

type Task struct {
	Urn     string `json:"urn"`
	Uri     string `json:"uri"`
	TaskUrn string `json:"taskUrn"`
	TaskUri string `json:"taskUri"`
}
