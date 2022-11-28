package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type setMethod func(ctx context.Context, key string, value interface{}, expiration time.Duration) error
type getMethod func(ctx context.Context, key string) (string, error)

type redisMock struct {
	MockSetItem setMethod
	MockGetItem getMethod
}

func (r redisMock) SetItem(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.MockSetItem(ctx, key, value, expiration)
}

func (r redisMock) GetItem(ctx context.Context, key string) (string, error) {
	return r.MockGetItem(ctx, key)
}

func TestRedis(t *testing.T) {
	redisCli := redisMock{}
	t.Run("Set item successfully", func(t *testing.T) {
		redisCli.MockSetItem = func(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
			return nil
		}

		err := redisCli.SetItem(context.Background(), "oneKey", "oneValue", 0)
		assert.Nil(t, err)
	})

	t.Run("Get item successfully", func(t *testing.T) {
		redisCli.MockGetItem = func(ctx context.Context, key string) (string, error) {
			return "item recovered", nil
		}

		resp, err := redisCli.GetItem(context.Background(), "oneKey")
		assert.Nil(t, err)
		assert.Equal(t, "item recovered", resp)
	})

}
