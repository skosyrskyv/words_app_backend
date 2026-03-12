package redis

import (
	"auth/config"
	"context"
	"log"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	db *redis.Client
}

func Init(ctx context.Context, cfg config.RedisConfig) *Redis {
	db := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		DB:       cfg.DB,
		Username: cfg.Username,
		Password: cfg.Password,
	})

	if err := db.Ping(ctx).Err(); err != nil {
		log.Fatal(err)
		return nil
	}
	slog.Info("Redis connected")

	return &Redis{
		db: db,
	}
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.db.Set(ctx, key, value, expiration).Err()
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.db.Get(ctx, key).Result()
}

func (r *Redis) Close() {
	r.db.Close()
	slog.Info("Redis connection closed")
}
