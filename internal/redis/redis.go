package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

var DefaultRedisConfig = RedisConfig{
	Host:     "redis",
	Port:     "6379",
	DB:       0,
	Username: "",
	Password: "",
}

// RedisConfig contains infos needed to open a postgres connection
type RedisConfig struct {
	Host     string
	Port     string
	DB       int
	Username string
	Password string
}

// GetAddr redis address
// e.g. localhost:6379
func (cfg *RedisConfig) GetAddr() string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}

// Returns new redis client
func NewRedisClient(cfg *RedisConfig) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.GetAddr(),
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       cfg.DB, // use default DB
	})

	return client
}
