package redis

import (
	"context"
	"fmt"

	"github.com/cenkalti/backoff/v4"

	"github.com/butlerhq/butler/internal/logger"
	"go.uber.org/zap"

	"github.com/go-redis/redis/v8"
)

var DefaultRedisConfig = Config{
	Host:     "redis",
	Port:     "6379",
	DB:       0,
	Username: "",
	Password: "",
}

// RedisConfig contains infos needed to open a postgres connection
type Config struct {
	Host     string `env:"HOST"`
	Port     string `env:"POSRT"`
	DB       int    `env:"DB"`
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
}

// GetAddr redis address
// e.g. localhost:6379
func (cfg *Config) GetAddr() string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}

type Redis struct {
	Config *Config
	Client *redis.Client
}

// Returns new redis client
func NewRedisClient(cfg *Config) *Redis {
	logger.Debug(context.Background(), "About to connect redis", zap.String("addr", cfg.GetAddr()))

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.GetAddr(),
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       cfg.DB, // use default DB
	})

	return &Redis{Config: cfg, Client: client}
}

// CheckConnection will try to connect to redis and will backoff if failed. It will try with a limit of attempt passed in parameters
func (rds *Redis) CheckConnection(maxAttempt uint64) error {
	bkoff := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), maxAttempt)
	return backoff.Retry(func() error {
		ctx := context.Background()
		_, err := rds.Client.Ping(ctx).Result()
		if err != nil {
			logger.Error(ctx, "Failed to reach redis.", zap.String("addr", rds.Config.GetAddr()), zap.Error(err))
			return err
		}
		logger.Info(ctx, "Successfully reach redis", zap.String("addr", rds.Config.GetAddr()))
		return nil
	}, bkoff)
}
