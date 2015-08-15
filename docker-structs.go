package dockerguard

type DockerInfo struct {
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
	ID         string  `json:"Id"`
	Hostname   string  `json:"Hostname"`
	Image      string  `json:"Image"`
	IPAddress  string  `json:"IPAddress"`
	MacAddress string  `json:"MacAddress"`
	SizeRootFs float64 `json:"SizeRootFs"`
	SizeRw     float64 `json:"SizeRw"`
	MemoryUsed float64 `json:"MemoryUsed"`
	State      struct {
		Dead       bool   `json:"Dead"`
		Error      string `json:"Error"`
		ExitCode   int    `json:"ExitCode"`
		OOMKilled  bool   `json:"OOMKilled"`
		Paused     bool   `json:"Paused"`
		Pid        int    `json:"Pid"`
		Restarting bool   `json:"Restarting"`
		Running    bool   `json:"Running"`
		StartedAt  string `json:"StartedAt"`
		FinishedAt string `json:"FinishedAt"`
	} `json:"State"`
}
