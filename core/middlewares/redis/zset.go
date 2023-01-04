package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// ZCard 获取有序集合中成员的数量
func ZCard(ctx context.Context, client *redis.Client, key string) (int64, error) {
	num, err := client.ZCard(ctx, key).Result()

	if err != nil && err != redis.Nil {
		return 0, err
	}

	return num, nil
}

// ZCount 用于统计有序集合中指定 score 值范围内的元素个数
func ZCount(ctx context.Context, client *redis.Client, key, min, max string) (int64, error) {
	num, err := client.ZCount(ctx, key, min, max).Result()
	if err != nil && err != redis.Nil {
		return 0, err
	}

	return num, nil
}

/*
ZAdd 用于将一个或多个成员添加到有序集合中，或者更新已存在成员的 score 值

	ZAdd(ctx, "myset", &redis.Z{
		Score:  float64(i),
			Member: fmt.Sprintf("member%d", i),
		})
*/
func ZAdd(ctx context.Context, client *redis.Client, key string, members ...*redis.Z) error {
	return client.ZAdd(ctx, key, members...).Err()
}

// ZRem 删除缓存
func ZRem(ctx context.Context, client *redis.Client, key string, members ...interface{}) (int64, error) {
	num, err := client.ZRem(ctx, key, members...).Result()
	if err != nil {
		zap.S().Error("cache ZRem fail:", key, err)
	}

	return num, err
}

// ZAddAndExpire 设置缓存&过期时间
func ZAddAndExpire(ctx context.Context, client *redis.Client, key string, expiration time.Duration, members ...*redis.Z) error {
	_, err := client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		client.ZAdd(ctx, key, members...)

		pipe.Expire(ctx, key, expiration)

		return nil
	})

	return err
}

// ZRange ZRangeByScore ZRangeByLex ZRank ZRevRange ZRevRank
