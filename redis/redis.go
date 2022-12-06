package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisIntf interface {
	SetItem(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	GetItem(ctx context.Context, key string) (string, error)
}

type redisCli struct {
	conn *redis.Client
}

func New(address, password string, db int) RedisIntf {
	rdb := redis.NewClient(&redis.Options{
		Addr:        address,
		Password:    password, // no password set
		DB:          db,       // use default DB
		ReadTimeout: 3 * time.Second,
		IdleTimeout: 5 * time.Minute,
		MaxRetries:  3,
	})

	return redisCli{
		conn: rdb,
	}
}

func (r redisCli) SetItem(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := r.conn.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r redisCli) GetItem(ctx context.Context, key string) (string, error) {
	val, err := r.conn.Get(ctx, key).Result()
	if err == redis.Nil {
		return val, err // Chequear esto
	} else if err != nil {
		return val, err
	}

	return val, nil
}
