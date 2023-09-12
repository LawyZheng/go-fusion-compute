package volume

type Volume struct {
	Urn              string `json:"urn"`
	URI              string `json:"uri"`
	UUID             string `json:"uuid"`
	Name             string `json:"name"`
	QuantityGB       int    `json:"quantityGB"`
	Status           string `json:"status"`
	StorageType      string `json:"storageType"`
	IsThin           bool   `json:"isThin"`
	Type             string `json:"type"`
	DatastoreUrn     string `json:"datastoreUrn"`
	DatastoreName    string `json:"datastoreName"`
	IndepDisk        bool   `json:"indepDisk"`
	PersistentDisk   bool   `json:"persistentDisk"`
	VolNameOnDev     string `json:"volNameOnDev"`
	VolProvisionSize int    `json:"volProvisionSize"`
	UserUsedSize     int    `json:"userUsedSize"`
	IsDiffVol        bool   `json:"isDiffVol"`
	VolType          int    `json:"volType"`
	MaxReadBytes     int    `json:"maxReadBytes"`
	MaxWriteBytes    int    `json:"maxWriteBytes"`
	MaxReadRequest   int    `json:"maxReadRequest"`
	MaxWriteRequest  int    `json:"maxWriteRequest"`
	TotalRWBytes     int    `json:"totalRWBytes"`
	TotalRWRequest   int    `json:"totalRWRequest"`
	PciType          string `json:"pciType"`
	SrcVolumeUrn     string `json:"srcVolumeUrn"`
	VolumeUseType    int    `json:"volumeUseType"`
	IoWeight         int    `json:"ioWeight"`
	SiocFlag         int    `json:"siocFlag"`
	VolumeURL        string `json:"volumeUrl"`
	VolInfoURL       string `json:"volInfoUrl"`
	DrExtParams      string `json:"drExtParams"`
	PvscsiSupport    int    `json:"pvscsiSupport"`
	StorageVersion   string `json:"storageVersion"`
	VolumeFormat     string `json:"volumeFormat"`
}
