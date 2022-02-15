package destinations

type Destination interface {
	Name() string
	BuildConfig(workspaceID string) interface{}
	BindAirbyteDefinition(destinationDefinitionID string)
	AirbyteDefinitionID() string
}
