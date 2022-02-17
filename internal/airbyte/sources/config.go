package sources

import (
	"errors"
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

var secret_keys = map[string]bool{"developer_token": true}

// SanitizeConfig remove secrets from airbyte config and put them in a separate map
func SanitizeConfig(cfgMap map[string]interface{}) (map[string]interface{}, map[string]interface{}) {
	secretsMap := make(map[string]interface{})
	for cfgKey, cfgValue := range cfgMap {
		if _, ok := secret_keys[cfgKey]; ok {
			secretsMap[cfgKey] = cfgValue
			delete(cfgMap, cfgKey)
		}
	}
	return cfgMap, secretsMap
}

// ValidateConfig  validate a json config against a json schema.
// It returns an error if it doesn't match the json schema
func ValidateConfig(schema string, payload interface{}) error {
	schemaLoader := gojsonschema.NewStringLoader(schema)
	documentLoader := gojsonschema.NewGoLoader(payload)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		return nil
	} else {
		errMsg := "Invalid Config : \n"
		for _, desc := range result.Errors() {
			errMsg = errMsg + fmt.Sprintf("- %s\n", desc)
		}
		return errors.New(errMsg)
	}
}
