package registry

type Registration struct {
	ServiceName      ServiceName
	ServiceURL       string
	RequiredServices []ServiceName
	ServiceUpdateURL string //服务注册中心告知服务所需的服务的URL
	HeartbeatURL     string
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradinService")
)

//每一条服务更新
type patchEntry struct {
	Name ServiceName
	URL  string
}

//更新的条目
type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
