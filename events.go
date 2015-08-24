package dockerguard

const (
	EventNotice = iota
	EventWarning
	EventCritical

	EventDiskSpaceLimitReached = iota
	EventMemorySpaceLimitReached
	EventContainerStarted
	EventContainerStopped
	EventContainerCreated
	EventContainerRemoved
	EventDiskIOOverload
	EventNetBandwithOverload
	EventCPUUsageOverload
)

type Event struct {
	Severity int
	Type     int
	Target   string
	Probe    string
	Data     string
}

func (e *Event) TypeToString() string {
	switch e.Type {
	case EventDiskSpaceLimitReached:
		return "DiskSpaceLimitReached"
	case EventMemorySpaceLimitReached:
		return "MemorySpaceLimitReached"
	case EventContainerStarted:
		return "ContainerStarted"
	case EventContainerStopped:
		return "ContainerStopped"
	case EventContainerRemoved:
		return "ContainerRemoved"
	case EventDiskIOOverload:
		return "DiskIOOverload"
	case EventNetBandwithOverload:
		return "NetBandwithOverload"
	case EventCPUUsageOverload:
		return "CPUUsageOverload"
	}

	return "UnknownType"
}
