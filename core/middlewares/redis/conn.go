package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

type Client struct {
	Sync   *redsync.Redsync
	Client *redis.Client
}

type Redis struct {
	Host         string `json:"host"`         // 地址
	Port         int    `json:"port"`         // 端口
	DB           int    `json:"db"`           // 库名
	User         string `json:"user"`         // 用户名
	Password     string `json:"password"`     // 密码
	PoolSize     int    `json:"poolSize"`     // 连接池最大socket连接数
	MinIdleConns int    `json:"minIdleConns"` // 闲置连接数量
}

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
