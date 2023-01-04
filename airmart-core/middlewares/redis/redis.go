package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
)

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
