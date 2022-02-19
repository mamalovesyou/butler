package sources

import (
	"database/sql/driver"

	"github.com/butlerhq/butler/api/services/octopus/v1"
)

type AuthScheme string

const (
	OAUTH2  AuthScheme = "OAUTH2"
	API_KEY AuthScheme = "API_KEY"
)

func (as *AuthScheme) Scan(value interface{}) error {
	*as = AuthScheme(value.(string))
	return nil
}

func (as AuthScheme) Value() (driver.Value, error) {
	return string(as), nil
}

func (as AuthScheme) String() string {
	return string(as)
}

func (as AuthScheme) ToPb() octopus.AuthType {
	value := octopus.AuthType_value[as.String()]
	return octopus.AuthType(value)
}
