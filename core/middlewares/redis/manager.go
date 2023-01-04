package redis

import (
	"crypto/rand"
	"math/big"
	rd "math/rand"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

type Manager struct {
	writeCache *redis.Pool
	readCaches []*redis.Pool
}

// RCache 随机返回一个读库.
func (cacheMgr *Manager) RCache() *redis.Pool {
	length := len(cacheMgr.readCaches)

	randomNum, _ := rand.Int(rand.Reader, big.NewInt(int64(length)))

	return cacheMgr.readCaches[randomNum.Int64()]
}

// WCache 返回唯一写库.
func (cacheMgr *Manager) WCache() *redis.Pool {
	return cacheMgr.writeCache
}

// InitManager 初始化缓存管理器.
func InitManager(conf *RedisConfig) (cacheMgr *Manager, err error) {
	var (
		wCache  *redis.Pool
		rCaches []*redis.Pool
		rCache  *redis.Pool
	)

	wCache, err = newCache(conf, conf.WCacheHost)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(conf.RCacheHost); i++ {
		rCache, err = newCache(conf, conf.RCacheHost[i])
		if err != nil {
			return nil, err
		}

		rCaches = append(rCaches, rCache)
	}

	rd.Seed(time.Now().Unix())

	cacheMgr = &Manager{
		writeCache: wCache,
		readCaches: rCaches,
	}

	return
}

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

type RedisConfig struct {
	WCacheHost         string
	RCacheHost         []string
	Pass               string
	DB                 int
	MaxIdle            int
	MaxActive          int
	IdleTimeout        int
	Wait               bool
	DialConnectTimeout int
	DialReadTimeout    int
	DialWriteTimeout   int
}
