package environment

import "strings"

// IsProductionEnv return true if environment variable ENV is set to prod | production
func IsProductionEnv(environment string) bool {
	value := strings.ToUpper(environment)
	return value == PROD || strings.ToUpper(environment) == PRODUCTION
}

// IsDevelopmentEnv return true if environment variable ENV is set to dev | development
func IsDevelopmentEnv(environment string) bool {
	value := strings.ToUpper(environment)
	return value == DEV || strings.ToUpper(environment) == DEVELOPMENT
}
