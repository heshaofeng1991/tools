package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// Gets kv缓存 批量获取
func Gets(ctx context.Context, client *redis.Client, key ...string) ([]interface{}, error) {
	value, err := client.MGet(ctx, key...).Result()

	if err != nil && err != redis.Nil {
		return value, err
	}

	return value, nil
}

// Get kv缓存 单条获取
func Get(ctx context.Context, client *redis.Client, key string) (string, bool) {
	value, err := client.Get(ctx, key).Result()
	if err == nil && len(value) > 0 {
		return value, true
	}

	return "", false
}

// Set 设置缓存  设置缓存强制要带过期时间 expire 1h 1s 1m 1d
func Set(ctx context.Context, client *redis.Client, key string, data string, expire string) error {
	duration, err := time.ParseDuration(expire)
	if err != nil {
		return err
	}

	return client.Set(ctx, key, data, duration).Err()
}

// Del 删除缓存
func Del(ctx context.Context, client *redis.Client, key string) (int64, error) {
	num, err := client.Del(ctx, key).Result()
	if err != nil {
		zap.S().Error("cache del fail:", key, err)
	}

	return num, err
}

func TTL(ctx context.Context, client *redis.Client, key string) error {
	err := client.TTL(ctx, key).Err()
	if err != nil {
		zap.S().Error("cache TTL fail:", key, err)
	}

	return err
}

// Expire 设置缓存过期时间
func Expire(ctx context.Context, client *redis.Client, key string, expiration time.Duration) (bool, error) {
	isExpired, err := client.Expire(ctx, key, expiration).Result()
	if err != nil {
		zap.S().Error("cache expire fail:", key, err)
	}

	return isExpired, err
}
