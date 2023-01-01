# Redis

## 数据结构
- redis连接池（实现了两种方式，conn.go,manager.go）
```go
conn.go

func New(conf *Redis) (*Client, error) {
	conn := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password, // 密码
		DB:       conf.DB,       // redis数据库index
	})
	ping := conn.Ping(context.Background())
	conn.AddHook(newJaegerHook())
	lock := redsync.New(goredis.NewPool(conn))

	r := &Client{}
	r.Sync = lock
	r.Client = conn

	return r, ping.Err()
}
```

```go
manager.go

// newCache redis连接池.
func newCache(conf *RedisConfig, host string) (*redis.Pool, error) {
	pool := &redis.Pool{
		MaxIdle:     conf.MaxIdle,
		MaxActive:   conf.MaxActive,
		IdleTimeout: time.Duration(conf.IdleTimeout) * time.Second,
		Wait:        conf.Wait,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", host,
				redis.DialPassword(conf.Pass),
				redis.DialDatabase(conf.DB),
				redis.DialConnectTimeout(time.Duration(conf.DialConnectTimeout)*time.Second),
				redis.DialReadTimeout(time.Duration(conf.DialReadTimeout)*time.Second),
				redis.DialWriteTimeout(time.Duration(conf.DialWriteTimeout)*time.Second))
			if err != nil {
				return nil, errors.Wrap(err, "")
			}

			return conn, nil
		},
	}

	conn := pool.Get()
	defer conn.Close()

	if _, err := conn.Do("ping"); err != nil {
		pool.Close()

		return nil, errors.Wrap(err, "")
	}

	return pool, nil
}
```
- hash
  - HGetAll 查询所有缓存数据
  - HGet 查询单条缓存数据
  - HSet 设置缓存数据
  - HDel 删除缓存
  - HLen 查询缓存长度
  - HExists 缓存是否存在
  - HSetAndExpire 设置缓存&过期时间
- plugin
    - redis Hook
    - BeforeProcess
    - AfterProcess
    - BeforeProcessPipeline
    - AfterProcessPipeline
- redis
  - 获取redis cmder
  - Lock
  - UnLock
  - 互斥锁(控制并发)
- string
  - Gets kv缓存 批量获取
  - Get kv缓存 单条获取
  - Set 设置缓存  设置缓存强制要带过期时间 expire 1h 1s 1m 1d
  - Del 删除缓存
  - TTL
  - Expire 设置缓存过期时间
- zset
  - ZCard 获取有序集合中成员的数量
  - ZCount 用于统计有序集合中指定 score 值范围内的元素个数
  - ZAdd 用于将一个或多个成员添加到有序集合中，或者更新已存在成员的 score 值
  - ZRem 删除缓存
  - ZAddAndExpire 设置缓存&过期时间
  - ZRange
  - ZRangeByScore
  - ZRangeByLex
  - ZRank
  - ZRevRange
  - ZRevRank
- list
- set
- options