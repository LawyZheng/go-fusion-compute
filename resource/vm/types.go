package vm

import (
	"context"
	"fmt"
	"strings"

	"github.com/lawyzheng/go-fusion-compute/client"
	"github.com/lawyzheng/go-fusion-compute/resource/site"
)

type Vm struct {
	Urn                string            `json:"urn,omitempty,omitempty"`
	Uri                string            `json:"uri,omitempty"`
	Uuid               string            `json:"uuid,omitempty"`
	Name               string            `json:"name,omitempty"`
	Arch               string            `json:"arch,omitempty"`
	Description        string            `json:"description,omitempty"`
	Group              string            `json:"group,omitempty"`
	Location           string            `json:"location,omitempty"`
	LocationName       string            `json:"locationName,omitempty"`
	HostUrn            string            `json:"hostUrn,omitempty"`
	Status             string            `json:"status,omitempty"`
	PvDriverStatus     string            `json:"pvDriverStatus,omitempty"`
	ToolInstallStatus  string            `json:"toolInstallStatus,omitempty"`
	CdRomStatus        string            `json:"cdRomStatus,omitempty"`
	IsTemplate         bool              `json:"isTemplate,omitempty"`
	IsLinkClone        bool              `json:"isLinkClone,omitempty"`
	IsBindingHost      bool              `json:"isBindingHost,omitempty"`
	IsMultiDiskSpeedup bool              `json:"isMultiDiskSpeedup,omitempty"`
	CreateTime         string            `json:"createTime,omitempty"`
	ToolsVersion       string            `json:"toolsVersion,omitempty"`
	HostName           string            `json:"hostName,omitempty"`
	ClusterName        string            `json:"clusterName,omitempty"`
	HugePage           string            `json:"hugePage,omitempty"`
	Idle               int               `json:"idle,omitempty"`
	VmType             int               `json:"vmType,omitempty"`
	DrStatus           int               `json:"drStatus,omitempty"`
	RpoStatus          int               `json:"rpoStatus,omitempty"`
	InitSyncStatus     int               `json:"initSyncStatus,omitempty"`
	ObjectPrivs        []string          `json:"objectPrivs,omitempty"`
	Params             map[string]string `json:"params,omitempty"`
	CustomProperties   map[string]string `json:"customProperties,omitempty"`
	RebootConfig       RebootConfig      `json:"vmRebootConfig,omitempty"`
	HAConfig           HAConfig          `json:"haConfig,omitempty"`
	VmConfig           Config            `json:"vmConfig,omitempty"`
	OsOption           OsOption          `json:"osOptions,omitempty"`
}

type RebootConfig struct {
	Cpu    Cpu    `json:"cpu,omitempty"`
	Memory Memory `json:"memory,omitempty"`
}

type HAConfig struct {
	HostFaultPolicy int `json:"hostFaultPolicy,omitempty"`
	VMFaultPolicy   int `json:"vmFaultPolicy,omitempty"`
}

type OsOption struct {
	OsType      string `json:"osType,omitempty"`
	OsVersion   int    `json:"osVersion,omitempty"`
	Hostname    string `json:"hostname,omitempty"`
	Password    string `json:"password,omitempty"`
	GuestOSName string `json:"guestOSName,omitempty"`
}

type Customization struct {
	OsType             string             `json:"osType,omitempty"`
	Hostname           string             `json:"hostname,omitempty"`
	IsUpdateVmPassword bool               `json:"isUpdateVmPassword,omitempty"`
	Password           string             `json:"password,omitempty"`
	NicSpecification   []NicSpecification `json:"nicSpecification,omitempty"`
}

type NicSpecification struct {
	SequenceNum int    `json:"sequenceNum,omitempty"`
	Ip          string `json:"ip,omitempty"`
	Netmask     string `json:"netmask,omitempty"`
	Gateway     string `json:"gateway,omitempty"`
	Setdns      string `json:"setdns,omitempty"`
	Adddns      string `json:"adddns,omitempty"`
}

type ListVmResponse struct {
	Total int  `json:"total,omitempty"`
	Vms   []Vm `json:"vms,omitempty"`
}

type CloneVmRequest struct {
	Name            string        `json:"name,omitempty"`
	Description     string        `json:"description,omitempty"`
	Group           string        `json:"group,omitempty"`
	Location        string        `json:"location,omitempty"`
	IsBindingHost   bool          `json:"isBindingHost,omitempty"`
	Config          Config        `json:"vmConfig,omitempty"`
	VmCustomization Customization `json:"vmCustomization,omitempty"`
}

type Config struct {
	Cpu          Cpu             `json:"cpu,omitempty"`
	Memory       Memory          `json:"memory,omitempty"`
	Disks        []Disk          `json:"disks,omitempty"`
	Nics         []Nic           `json:"nics,omitempty"`
	Property     Property        `json:"properties"`
	USB          []USBController `json:"usb"`
	NUMANodes    int             `json:"numaNodes"`
	GraphicsCard struct {
		Size int    `json:"size"`
		Type string `json:"type"`
	} `json:"graphicsCard"`
}

func (c *Config) ClearConfig() {
	for i := range c.Nics {
		c.Nics[i].ClearPrivate()
	}

	for i := range c.Disks {
		c.Disks[i].ClearPrivate()
	}
}

type USBController struct {
	ControllerType string   `json:"controllerType"`
	Device         []string `json:"device"`
}

type Cpu struct {
	Quantity        int    `json:"quantity,omitempty"`
	Reservation     int    `json:"reservation,omitempty"`
	Weight          int    `json:"weight,omitempty"`
	Limit           int    `json:"limit,omitempty"`
	CoresPerSocket  int    `json:"coresPerSocket,omitempty"`
	CPUHotPlug      int    `json:"cpuHotPlug,omitempty"`
	CPUPolicy       string `json:"cpuPolicy,omitempty"`
	CPUThreadPolicy string `json:"cpuThreadPolicy,omitempty"`
	CPUBindType     string `json:"cpuBindType,omitempty"`
	NumaBinds       []struct {
		NodeId int    `json:"nodeId,omitempty"`
		VCPUs  string `json:"vcpus,omitempty"`
	} `json:"numaBinds,omitempty"`
}

type Memory struct {
	QuantityMB  int    `json:"quantityMB,omitempty"`
	Reservation int    `json:"reservation,omitempty"`
	Weight      int    `json:"weight,omitempty"`
	Limit       int    `json:"limit,omitempty"`
	MemHotPlug  int    `json:"memHotPlug,omitempty"`
	HugePage    string `json:"hugePage,omitempty	"`
}

type Disk struct {
	SequenceNum     int    `json:"sequenceNum,omitempty"`
	QuantityGB      int    `json:"quantityGB,omitempty"`
	IsDataCopy      bool   `json:"isDataCopy,omitempty"`
	DatastoreUrn    string `json:"datastoreUrn,omitempty"`
	IsThin          bool   `json:"isThin,omitempty"`
	VolumeUrn       string `json:"volumeUrn,omitempty"`
	VolumeUUID      string `json:"volumeUuid,omitempty"`
	IndepDisk       bool   `json:"indepDisk,omitempty"`
	PersistentDisk  bool   `json:"persistentDisk,omitempty"`
	StorageType     string `json:"storageType,omitempty"`
	VolType         int    `json:"volType,omitempty"`
	PciType         string `json:"pciType,omitempty"`
	MaxReadBytes    int    `json:"maxReadBytes,omitempty"`
	MaxWriteBytes   int    `json:"maxWriteBytes,omitempty"`
	MaxReadRequest  int    `json:"maxReadRequest,omitempty"`
	MaxWriteRequest int    `json:"maxWriteRequest,omitempty"`
	TotalRWBytes    int    `json:"totalRWBytes,omitempty"`
	TotalRWRequest  int    `json:"totalRWRequest,omitempty"`
	IoWeight        int    `json:"ioWeight,omitempty"`
	DevName         string `json:"devName,omitempty"`
	Type            string `json:"type,omitempty"`
	DiskName        string `json:"diskName,omitempty"`
	VolumeURL       string `json:"volumeUrl,omitempty"`
	DatastoreName   string `json:"datastoreName,omitempty"`
	BootOrder       int    `json:"bootOrder,omitempty"`
	IoMode          string `json:"ioMode,omitempty"`
	SystemVolume    bool   `json:"systemVolume,omitempty"`
	VolumeFormat    string `json:"volumeFormat,omitempty"`
	Thin            bool   `json:"thin,omitempty"`
	DataCopy        bool   `json:"dataCopy,omitempty"`
}

func (d *Disk) IsRBD() bool {
	return d.StorageType == "FusionOneStorage"
}

func (d *Disk) ClearPrivate() {
	d.VolumeUUID = ""
	d.VolumeUrn = ""
	d.VolumeURL = ""
	d.DiskName = ""
}

type Nic struct {
	Name          string `json:"name,omitempty"`
	PortGroupUrn  string `json:"portGroupUrn,omitempty"`
	PortGroupName string `json:"portGroupName,omitempty"`
	Mac           string `json:"mac,omitempty"`
	IP            string `json:"ip,omitempty"`
	Urn           string `json:"urn,omitempty"`
	URI           string `json:"uri,omitempty"`
	Status        string `json:"status,omitempty"`
	NicConfig     struct {
		Vringbuf int `json:"vringbuf,omitempty"`
		Queues   int `json:"queues,omitempty"`
	} `json:"nicConfig,omitempty"`
	IPList              string   `json:"ipList,omitempty"`
	SequenceNum         int      `json:"sequenceNum,omitempty"`
	Ips6                []string `json:"ips6,omitempty"`
	VirtIo              int      `json:"virtIo,omitempty"`
	NicType             int      `json:"nicType,omitempty"`
	PortID              string   `json:"portId,omitempty"`
	BootOrder           int      `json:"bootOrder,omitempty"`
	EnableSecurityGroup bool     `json:"enableSecurityGroup,omitempty"`
	SecurityGroupName   string   `json:"securityGroupName,omitempty"`
}

func (n *Nic) ClearPrivate() {
	n.Mac = ""
	n.IP = ""
	n.IPList = ""
	n.Ips6 = []string{}
}

type Property struct {
	BootFirmware       string `json:"bootFirmware,omitempty"`
	VMVncKeymapSetting int    `json:"vmVncKeymapSetting,omitempty"`
	CdRomBootOrder     int    `json:"cdRomBootOrder,omitempty"`
	Consolelog         int    `json:"consolelog,omitempty"`
	Realtime           bool   `json:"realtime,omitempty"`
	BootFirmwareTime   int    `json:"bootFirmwareTime,omitempty"`
	GpuShareType       string `json:"gpuShareType,omitempty"`
	IsReserveResource  bool   `json:"isReserveResource,omitempty"`
	IsHpet             bool   `json:"isHpet,omitempty"`
	BootOption         string `json:"bootOption,omitempty"`
	VMFaultProcess     string `json:"vmFaultProcess,omitempty"`
	ClockMode          string `json:"clockMode,omitempty"`
	IsEnableMemVol     bool   `json:"isEnableMemVol,omitempty"`
	IsEnableFt         bool   `json:"isEnableFt,omitempty"`
	IsAutoUpgrade      bool   `json:"isAutoUpgrade,omitempty"`
	AttachType         bool   `json:"attachType,omitempty"`
	ReoverByHost       bool   `json:"reoverByHost,omitempty"`
	EvsAffinity        bool   `json:"evsAffinity,omitempty"`
	IsAutoAdjustNuma   bool   `json:"isAutoAdjustNuma,omitempty"`
	SystemArchitecture string `json:"systemArchitecture,omitempty"`
	NumaAffinity       []int  `json:"numaAffinity,omitempty"`
	EmulatorPin        []int  `json:"emulatorPin,omitempty"`
}

type CloneVmResponse struct {
	Urn     string `json:"urn,omitempty"`
	Uri     string `json:"uri,omitempty"`
	TaskUrn string `json:"taskUrn,omitempty"`
	TaskUri string `json:"taskUri,omitempty"`
}

type DeleteVmResponse struct {
	TaskUrn string `json:"taskUrn,omitempty"`
	TaskUri string `json:"taskUri,omitempty"`
}

type ImportTemplateRequest struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Location    string   `json:"location,omitempty"`
	VmConfig    Config   `json:"vmConfig,omitempty"`
	OsOptions   OsOption `json:"osOptions,omitempty"`
	Url         string   `json:"url,omitempty"`
	Protocol    string   `json:"protocol,omitempty"`
	IsTemplate  bool     `json:"isTemplate,omitempty"`
}

type ImportTemplateResponse struct {
	TaskUrn string `json:"taskUrn,omitempty"`
	TaskUri string `json:"taskUri,omitempty"`
}

type CreateVMRequest struct {
	Name          string       `json:"name,omitempty"`
	Description   string       `json:"description,omitempty"`
	Group         string       `json:"group,omitempty"`
	Location      string       `json:"location,omitempty"`
	IsBindingHost bool         `json:"isBindingHost,omitempty"`
	AutoBoot      bool         `json:"autoBoot"`
	VMConfig      Config       `json:"vmConfig,omitempty"`
	OsOptions     OsOption     `json:"osOptions,omitempty"`
	RebootConfig  RebootConfig `json:"vmRebootConfig,omitempty"`
	HAConfig      HAConfig     `json:"haConfig,omitempty"`
}

type Task struct {
	Urn     string `json:"urn"`
	Uri     string `json:"uri"`
	TaskUrn string `json:"taskUrn"`
	TaskUri string `json:"taskUri"`
}

func searchVmById(ctx context.Context, c client.FusionComputeClient, vmM Manager, vmId string) (string, error) {
	siteM := site.NewManager(c)
	sites, err := siteM.ListSite(ctx)
	if err != nil {
		return "", err
	}
	for _, s := range sites {
		vms, err := vmM.ListVm(ctx, s.Uri, false)
		if err != nil {
			return "", err
		}

		for _, v := range vms {
			l := strings.Split(v.Urn, ":")
			if len(l) == 0 {
				continue
			}
			if l[len(l)-1] == vmId {
				return v.Uri, nil
			}
		}
	}
	return "", fmt.Errorf("vm[%s] not found", vmId)
}
