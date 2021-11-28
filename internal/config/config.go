package config

type Config struct {
	Environment string
	Log         *LoggerConfig
}

type LoggerConfig struct {
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}
