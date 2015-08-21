package dockerguard

const (
	DiskSpaceLimitReached   = iota
	MemorySpaceLimitReached = iota
	ContainerStarted        = iota
	ContainerStopped        = iota
	ContainerRemoved        = iota
	DiskIOOverload          = iota
	NetBandwithOverload     = iota
	CPUUsageOverload        = iota
)

type Event struct {
	Type   int
	Target string
	Data   string
}
