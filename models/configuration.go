package models

// Configuration of Overseer
// type Configuration struct {
// 	// 进程名称
// 	Name string
// 	// 地域
// 	Region string
// 	// 可用区
// 	AvailabilityZone string
// 	// 网络区（DMZ/Intranet）
// 	NetworkZone string
// 	// 云平台代号
// 	CloudPlatform string
// 	// 私有IP
// 	PrivateIP string
// 	// 公网IP
// 	PublicIP string
// 	// 存储类型
// 	VolumeType string
// 	// 存储读写路径
// 	VolumePath string
// 	// 加解密key
// 	Key string
// 	// 网络探测点
// 	Probes []EndPoint
// }

// Configuration of Overseer
type Configuration struct {
	Cloud            string   `json:"cloud"`
	Region           string   `json:"region"`
	AvailabilityZone string   `json:"availabilityZone"`
	NetworkZone      string   `json:"networkZone"`
	PrivateIP        string   `json:"privateIp"`
	PublicIP         string   `json:"publicIp"`
	Port             int      `json:"port"`
	Key              string   `json:"key"`
	Overseer         Overseer `json:"overseer"`
}

// Probe target
type Probe struct {
	Endpoint     string `json:"endpoint"`
	EndpointName string `json:"endpointName"`
}

// Volume of storage check items
type Volume struct {
	Path string `json:"path"`
	Type string `json:"type"`
}

// Self reporter
type Self struct {
	Interval   string `json:"interval"`
	APIEnabled bool   `json:"apiEnabled"`
}

// Overseer check items
type Overseer struct {
	Probes  []Probe  `json:"probes"`
	Volumes []Volume `json:"volumes"`
	Self    Self     `json:"self"`
}

// ToPing from Configuration
func (c *Configuration) ToPing() Ping {
	return Ping{
		Cloud:            c.Cloud,
		Region:           c.Region,
		AvailabilityZone: c.AvailabilityZone,
		NetworkZone:      c.NetworkZone,
	}
}
