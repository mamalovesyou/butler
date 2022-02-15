package connections

import (
	"encoding/json"

	"github.com/butlerhq/butler/internal/airbyte/sources"

	"github.com/butlerhq/airbyte-client-go/airbyte"
)

type ScheduleConfig struct {
	Units    int8   `json:"units"`
	TimeUnit string `json:"timeUnit"`
}

type ConnectionConfig struct {
	SourceID            string               `json:"sourceId"`
	DestinationID       string               `json:"destinationId"`
	Status              string               `json:"status"`
	Schedule            *ScheduleConfig      `json:"schedule"`
	NamespaceFormat     string               `json:"namespaceFormat"`
	NamespaceDefinition string               `json:"namespaceDefinition"`
	SyncCatalog         *sources.SyncCatalog `json:"syncCatalog"`
	Prefix              string               `json:"prefix"`
}

func NewConnectionConfig(sourceID, destinationID, jsonCatalog string) *ConnectionConfig {

	catalog := &sources.SyncCatalog{}
	if err := json.Unmarshal([]byte(jsonCatalog), catalog); err != nil {
		panic(err)
	}

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

func (cfg *ConnectionConfig) ToAirbyteCreateConnectionRequestBody() *airbyte.CreateConnectionJSONRequestBody {
	bytes, err := json.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	result := &airbyte.CreateConnectionJSONRequestBody{}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		panic(err)
	}
	return result
}
