package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
)

type RedisIntf interface {
	SetItem(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	GetItem(ctx context.Context, key string) (string, error)
	InsertElementInList(ctx context.Context, key string) (int64, error)
	GetElementsInList(ctx context.Context, key string, start int64, stop int64) ([]string, error)
}

type redisCli struct {
	conn *redis.Client
}

func New(host, port, password string, db int) (RedisIntf, error) {
	address := fmt.Sprintf("%s:%s", host, port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, errors.New("error creating redis client: " + err.Error())
	}

	return redisCli{
		conn: rdb,
	}, nil
}

func (r redisCli) SetItem(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := r.conn.Set(ctx, toLowerCase(key), value, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r redisCli) GetItem(ctx context.Context, key string) (string, error) {
	val, err := r.conn.Get(ctx, toLowerCase(key)).Result()
	if err == redis.Nil {
		return val, errors.New("item not found")
	} else if err != nil {
		return val, err
	}

	return val, nil
}

func (r redisCli) InsertElementInList(ctx context.Context, key string) (int64, error) {
	val, err := r.conn.LPush(ctx, toLowerCase(key)).Result()
	if err == redis.Nil {
		return val, errors.New("error inserting element in list")
	} else if err != nil {
		return val, err
	}

	return val, nil
}

func (r redisCli) GetElementsInList(ctx context.Context, key string, start int64, stop int64) ([]string, error) {
	val, err := r.conn.LRange(ctx, toLowerCase(key), start, stop).Result()
	if err == redis.Nil {
		return val, errors.New("error getting element in list")
	} else if err != nil {
		return val, err
	}

	return val, nil
}

func toLowerCase(key string) string {
	return strings.ToLower(key)
}
