package dockerguard

type DockerInfos struct {
	// Infos from docker api /version
	APIVersion    string `json:"ApiVersion"`
	Arch          string `json:"Arch"`
	Experimental  bool   `json:"Experimental"`
	GitCommit     string `json:"GitCommit"`
	GoVersion     string `json:"GoVersion"`
	KernelVersion string `json:"KernelVersion"`
	Os            string `json:"Os"`
	Version       string `json:"Version"`

	// Infos from docker api /info
	ID              string `json:"ID"`
	Name            string `json:"Name"`
	Containers      int    `json:"Containers"`
	Images          int    `json:"Images"`
	Driver          string `json:"Driver"`
	SystemTime      string `json:"SystemTime"`
	OperatingSystem string `json:"OperatingSystem"`
	NCPU            int    `json:"NCPU"`
	MemTotal        int    `json:"MemTotal"`
}

type Container struct {
	ID            string  `json:"Id"`
	Hostname      string  `json:"Hostname"`
	Probe         string  `json:"Probe"`
	Image         string  `json:"Image"`
	IPAddress     string  `json:"IPAddress"`
	MacAddress    string  `json:"MacAddress"`
	SizeRootFs    float64 `json:"SizeRootFs"`
	SizeRw        float64 `json:"SizeRw"`
	MemoryUsed    float64 `json:"MemoryUsed"`
	NetBandwithRX float64 `json:"NetBandwithRX"`
	NetBandwithTX float64 `json:"NetBandwithTX"`
	CPUUsage      float64 `json:"CPUUsage"`
	Running       bool    `json:"Running"`
	Time          float64 `json:"Time"`
}

type SimpleContainer struct {
	ID         string `json:"Id"`
	Hostname   string `json:"Hostname"`
	Image      string `json:"Image"`
	IPAddress  string `json:"IPAddress"`
	MacAddress string `json:"MacAddress"`
}

type ProbeInfos struct {
	Name            string      `json:"Name"`
	Running         bool        `json:"Running"`
	LoadAvg         string      `json:"LoadAvg"`
	MemoryTotal     float64     `json:"MemoryTotal"`
	MemoryAvailable float64     `json:"MemoryAvailable"`
	DiskTotal       float64     `json:"DiskTotal"`
	DiskAvailable   float64     `json:"DiskAvailable"`
	Containers      []Container `json:"Containers"`
}
