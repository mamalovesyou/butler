package sources

import (
	"reflect"
	"testing"
)

var nonSanitzedConfig = map[string]interface{}{
	"developer_token": "qwerty",
	"non_secret_key":  124,
}

func TestSanitizeConfig(t *testing.T) {
	var tests = []struct {
		input           map[string]interface{}
		expectedConfig  map[string]interface{}
		expectedSecrets map[string]interface{}
	}{
		{
			input:           map[string]interface{}{"developer_token": "qwerty", "non_secret_key": 124},
			expectedConfig:  map[string]interface{}{"non_secret_key": 124},
			expectedSecrets: map[string]interface{}{"developer_token": "qwerty"},
		},
	}

	for _, c := range tests {
		cfgMap, secretMap := SanitizeConfig(c.input)
		if eq := reflect.DeepEqual(cfgMap, c.expectedConfig); !eq {
			t.Errorf("Wrong config map received, got=%+v, expected=%+v", cfgMap, c.expectedConfig)
		}
		if eq := reflect.DeepEqual(secretMap, c.expectedSecrets); !eq {
			t.Errorf("Wrong secrets map received, got=%+v, expected=%+v", secretMap, c.expectedSecrets)
		}
	}
}
