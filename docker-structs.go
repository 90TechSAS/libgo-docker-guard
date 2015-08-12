package dockerguard

type DockerInfo struct {
	Containers         int         `json:"Containers"`
	Debug              int         `json:"Debug"`
	DockerRootDir      string      `json:"DockerRootDir"`
	Driver             string      `json:"Driver"`
	DriverStatus       [][]string  `json:"DriverStatus"`
	ExecutionDriver    string      `json:"ExecutionDriver"`
	ID                 string      `json:"ID"`
	IPv4Forwarding     int         `json:"IPv4Forwarding"`
	Images             int         `json:"Images"`
	IndexServerAddress string      `json:"IndexServerAddress"`
	InitPath           string      `json:"InitPath"`
	InitSha1           string      `json:"InitSha1"`
	KernelVersion      string      `json:"KernelVersion"`
	Labels             interface{} `json:"Labels"`
	MemTotal           int         `json:"MemTotal"`
	MemoryLimit        int         `json:"MemoryLimit"`
	NCPU               int         `json:"NCPU"`
	NEventsListener    int         `json:"NEventsListener"`
	NFd                int         `json:"NFd"`
	NGoroutines        int         `json:"NGoroutines"`
	Name               string      `json:"Name"`
	OperatingSystem    string      `json:"OperatingSystem"`
	RegistryConfig     struct {
		IndexConfigs struct {
			Docker_io struct {
				Mirrors  interface{} `json:"Mirrors"`
				Name     string      `json:"Name"`
				Official bool        `json:"Official"`
				Secure   bool        `json:"Secure"`
			} `json:"docker.io"`
		} `json:"IndexConfigs"`
		InsecureRegistryCIDRs []string `json:"InsecureRegistryCIDRs"`
	} `json:"RegistryConfig"`
	SwapLimit  int    `json:"SwapLimit"`
	SystemTime string `json:"SystemTime"`
}

type DockerVersion struct {
	APIVersion    string `json:"ApiVersion"`
	Arch          string `json:"Arch"`
	Experimental  bool   `json:"Experimental"`
	GitCommit     string `json:"GitCommit"`
	GoVersion     string `json:"GoVersion"`
	KernelVersion string `json:"KernelVersion"`
	Os            string `json:"Os"`
	Version       string `json:"Version"`
}
