package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// HGetAll 查询所有缓存数据
func HGetAll(ctx context.Context, client *redis.Client, key string) (map[string]string, int64, error) {
	value, err := client.HGetAll(ctx, key).Result()

	if err != nil && err != redis.Nil {
		return value, 0, err
	}

	return value, int64(len(value)), nil
}

// HGet 查询单条缓存数据
func HGet(ctx context.Context, client *redis.Client, key, field string) (string, error) {
	value, err := client.HGet(ctx, key, field).Result()
	if err == nil && len(value) > 0 {
		return value, nil
	}

	return "", fmt.Errorf("hget failed")
}

// HSet 设置缓存数据
func HSet(ctx context.Context, client *redis.Client, key string, value ...interface{}) error {
	return client.HSet(ctx, key, value).Err()
}

// HDel 删除缓存
func HDel(ctx context.Context, client *redis.Client, key string, field ...string) (int64, error) {
	num, err := client.HDel(ctx, key, field...).Result()
	if err != nil {
		zap.S().Error("cache hdel fail:", key, err)
	}

	return num, err
}

// HLen 查询缓存长度
func HLen(ctx context.Context, client *redis.Client, key string) (int64, error) {
	num, err := client.HLen(ctx, key).Result()
	if err != nil {
		zap.S().Error("cache hlen fail:", key, err)
	}

	return num, err
}

// HExists 缓存是否存在
func HExists(ctx context.Context, client *redis.Client, key, filed string) (bool, error) {
	num, err := client.HExists(ctx, key, filed).Result()
	if err != nil {
		zap.S().Error("cache hexists fail:", key, err)
	}

	return num, err
}

// HSetAndExpire 设置缓存&过期时间
func HSetAndExpire(ctx context.Context, client *redis.Client, key string, expiration time.Duration, value ...string) error {
	_, err := client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.HSet(ctx, key, value)

		pipe.Expire(ctx, key, expiration)

		return nil
	})

	return err
}
