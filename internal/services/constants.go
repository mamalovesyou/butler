package services

const (
	GatewayServiceName   = "gateway"
	AuthServiceName      = "auth"
	WorkspaceServiceName = "workspace"
)

var ServiceWithPostgresDBList = []string{
	AuthServiceName,
	WorkspaceServiceName,
}
