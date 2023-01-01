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
```go
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
```
- plugin
    - redis Hook
    - BeforeProcess
    - AfterProcess
    - BeforeProcessPipeline
    - AfterProcessPipeline
```go
func (j *jaegerHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	return trace.ContextWithSpan(otel.Tracer(cmd.Name()).Start(ctx, optRedis+cmd.Name())), nil
}

func (j *jaegerHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	span := trace.SpanFromContext(ctx)
	if span.IsRecording() {
		str := fmt.Sprintf("%+v", cmd)
		log.New(ctx).Info(str)
		span.SetAttributes(attribute.String(CmdResult, str))
		span.End()
	}
	return nil
}

func (j *jaegerHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	return trace.ContextWithSpan(otel.Tracer("pipeline").Start(ctx, optRedis+"pipeline")), nil
}

func (j *jaegerHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	span := trace.SpanFromContext(ctx)
	if span.IsRecording() {
		defer span.End()
		str := fmt.Sprintf("%+v", cmds)
		log.New(ctx).Info(str)
		span.SetAttributes(attribute.String(CmdResult, str))
		attribute.String(CmdResult, str)
	}
	return nil
}
```
- redis
  - 获取redis cmder
  - Lock
  - UnLock
  - 互斥锁(控制并发)
```go
func (r *Client) Get(ctx context.Context, key string) *redis.StringCmd {
	cmd := r.Client.Get(ctx, key)
	defer r.Client.Close()

	return cmd
}

func (r *Client) Lock(name string, options ...redsync.Option) (mutex *redsync.Mutex, err error) {
	mutex = r.Sync.NewMutex(name, options...)
	err = mutex.Lock()

	return
}

// LockContext 互斥锁 控制并发
func LockContext(ctx context.Context, sync *redsync.Redsync, name string, t time.Duration, option ...redsync.Option) (mutex *redsync.Mutex, err error) {
	if len(option) == 0 {
		option = append(option, redsync.WithExpiry(t))
	}

	mux := sync.NewMutex(name, option...)

	return mux, mux.LockContext(ctx)
}
```
- string
  - Gets kv缓存 批量获取
  - Get kv缓存 单条获取
  - Set 设置缓存  设置缓存强制要带过期时间 expire 1h 1s 1m 1d
  - Del 删除缓存
  - TTL
  - Expire 设置缓存过期时间
```go
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
```
- zset
  - ZCard 获取有序集合中成员的数量
  - ZCount 用于统计有序集合中指定 score 值范围内的元素个数
  - ZAdd 用于将一个或多个成员添加到有序集合中，或者更新已存在成员的 score 值
  - ZRem 删除缓存
  - ZAddAndExpire 设置缓存&过期时间
  - ZRange(待补充完善)
  - ZRangeByScore(待补充完善)
  - ZRangeByLex(待补充完善)
  - ZRank(待补充完善)
  - ZRevRange(待补充完善)
  - ZRevRank(待补充完善)
```go
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
```
- list
- set
- options
```go
type option func(*options)

type options struct {}

func defaultConfig() *options {
	return &options{}
}

func NewOptions(opts ...option) *options {
	options := defaultConfig()
	for _, opt := range opts {
		opt(options)
	}
	
	return options
}
```