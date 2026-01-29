package redis

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

const (
	MaxRetries        = 3
	MinRetriesBackoff = 300 * time.Millisecond // in milliseconds
	MaxRetriesBackoff = 512 * time.Millisecond // in milliseconds
	MinIdleConns      = 10
	MaxIdleConns      = 100
	ReadTimeout       = 5 * time.Second // in seconds
	WriteTimeout      = 5 * time.Second // in seconds
	DialTimeout       = 5 * time.Second // in seconds
)

func NewRedisClient(cfg *RedisOptions) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:            fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:        cfg.Password,
		DB:              cfg.DB,
		PoolSize:        cfg.PoolSize,
		MaxRetries:      MaxRetries,
		MinRetryBackoff: MinRetriesBackoff,
		MaxRetryBackoff: MaxRetriesBackoff,
		MinIdleConns:    MinIdleConns,
		MaxIdleConns:    MaxIdleConns,
		ReadTimeout:     ReadTimeout,
		WriteTimeout:    WriteTimeout,
		DialTimeout:     DialTimeout,
	})

	if cfg.EnableTracing {
		_ = redisotel.InstrumentTracing(rdb)
	}

	return rdb
}
