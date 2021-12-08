package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

// RedisConfig contains infos needed to open a postgres connection
type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	DB           int    `mapstructure:"db`
	MinIdleConns int
	PoolSize     int
	Password     string
}

// GetAddr redis address
// e.g. localhost:6379
func (cfg *RedisConfig) GetAddr() string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}

// Returns new redis client
func NewRedisClient(cfg *RedisConfig) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr: cfg.GetAddr(),
		// MinIdleConns: cfg.MinIdleConns,
		// PoolSize:     cfg.PoolSize,
		// Password:     cfg.Password, // no password set
		DB: cfg.DB, // use default DB
	})

	return client
}
