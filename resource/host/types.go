package host

type listHostsResponse struct {
	Hosts []Host `json:"hosts"`
	Total int    `json:"total"`
}

type Host struct {
	Urn                    string   `json:"urn"`
	Uri                    string   `json:"uri"`
	Uuid                   string   `json:"uuid"`
	Name                   string   `json:"name"`
	Description            string   `json:"description"`
	IP                     string   `json:"ip"`
	BmcIP                  string   `json:"bmcIp"`
	BmcUserName            string   `json:"bmcUserName"`
	ClusterUrn             string   `json:"clusterUrn"`
	ClusterName            string   `json:"clusterName"`
	Status                 string   `json:"status"`
	IsMaintaining          bool     `json:"isMaintaining"`
	MultiPathMode          string   `json:"multiPathMode"`
	MemQuantityMB          int64    `json:"memQuantityMB"`
	CpuQuantity            int      `json:"cpuQuantity"`
	NicQuantity            int      `json:"nicQuantity"`
	CpuMHz                 int      `json:"cpuMHz"`
	AttachedISOVMs         []string `json:"attachedISOVMs"`
	ComputeResourceStatics string   `json:"computeResourceStatics"`
	NtpIP1                 string   `json:"ntpIp1"`
	NtpCycle               int      `json:"ntpCycle"`
	PhysicalCPUQuantity    int      `json:"physicalCpuQuantity"`
	GpuCapacity            int      `json:"gpuCapacity"`
	GpuCapacityReboot      int      `json:"gpuCapacityReboot"`
	GdvmMemoryReboot       int      `json:"gdvmMemoryReboot"`
	GsvmMemoryReboot       int      `json:"gsvmMemoryReboot"`
	MaxImcSetting          string   `json:"maxImcSetting"`
	IsFailOverHost         bool     `json:"isFailOverHost"`
	ClusterEnableIOTailor  bool     `json:"clusterEnableIOTailor"`
	HostRealName           string   `json:"hostRealName"`
	CPUResource            struct {
		TotalSizeMHz     int    `json:"totalSizeMHz"`
		AllocatedSizeMHz int    `json:"allocatedSizeMHz"`
		ManageCPUs       string `json:"manageCPUs"`
		EmulatorCPUs     string `json:"emulatorCPUs"`
		ReserveSizeMHz   int    `json:"reserveSizeMHz"`
	} `json:"cpuResource"`
	MemResource struct {
		TotalSizeMB        int `json:"totalSizeMB"`
		AllocatedSizeMB    int `json:"allocatedSizeMB"`
		RealtimeUsedSizeMB int `json:"realtimeUsedSizeMB"`
		ReserveSizeMB      int `json:"reserveSizeMB"`
	} `json:"memResource"`
	MemMuxRatio string `json:"memMuxRatio"`
	CPUMuxRatio string `json:"cpuMuxRatio"`
}
