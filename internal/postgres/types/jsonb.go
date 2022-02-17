package types

import (
	"database/sql/driver"
	"encoding/json"

	structpb "google.golang.org/protobuf/types/known/structpb"
)

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}

func (j *JSONB) ToPbStruct() *structpb.Struct {
	m := map[string]interface{}(*j)
	result, err := structpb.NewStruct(m)
	if err != nil {
		panic(err)
	}
	return result
}

func (j *JSONB) AsBytes() []byte {
	valueBytes, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	return valueBytes
}
