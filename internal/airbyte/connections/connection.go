package airbyte

type SyncCatalog struct {
	Streams []map[string]interface{} `json:"streams"`
}

type ScheduleConfig struct {
	Units    int8   `json:"units"`
	TimeUnit string `json:"timeUnit"`
}

type ConnectionConfig struct {
	SourceID            string          `json:"sourceId"`
	DestinationID       string          `json:"destinationId"`
	Status              string          `json:"status"`
	Schedule            *ScheduleConfig `json:"schedule"`
	NamespaceFormat     string          `json:"namespaceFormat"`
	NamespaceDefinition string          `json:"namespaceDefinition"`
	SyncCatalog         *SyncCatalog    `json:"syncCatalog"`
	Prefix              string          `json:"prefix"`
}

func NewConnectionConfig(sourceID, destinationID string, catalog *SyncCatalog) *ConnectionConfig {
	return &ConnectionConfig{
		SourceID:            sourceID,
		DestinationID:       destinationID,
		SyncCatalog:         catalog,
		Status:              "active",
		NamespaceFormat:     "${SOURCE_NAMESPACE}",
		NamespaceDefinition: "source",
		Schedule: &ScheduleConfig{
			TimeUnit: "hours",
			Units:    3,
		},
	}
}
