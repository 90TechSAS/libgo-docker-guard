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

func (e *Event) TypeToString() string {
	switch e.Type {
	case DiskSpaceLimitReached:
		return "DiskSpaceLimitReached"
	case MemorySpaceLimitReached:
		return "MemorySpaceLimitReached"
	case ContainerStarted:
		return "ContainerStarted"
	case ContainerStopped:
		return "ContainerStopped"
	case ContainerRemoved:
		return "ContainerRemoved"
	case DiskIOOverload:
		return "DiskIOOverload"
	case NetBandwithOverload:
		return "NetBandwithOverload"
	case CPUUsageOverload:
		return "CPUUsageOverload"
	}

	return "UnknownType"
}
