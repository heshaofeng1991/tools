package redis

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/go-redsync/redsync/v4"
)

var client *Client

func init() {
	conf := &Redis{
		Host: "192.168.8.243",
		Port: 6379,
		DB:   1,
	}
	var err error
	client, err = New(conf)
	if err != nil {
		panic(err)
	}
}

func TestClient_Lock(t *testing.T) {
	wait := sync.WaitGroup{}
	wait.Add(1)
	go func() {
		mutex, err := client.Lock("test")
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 10)
		mutex.Unlock()
		fmt.Println("解锁了")
		wait.Done()
	}()
	time.Sleep(time.Second)
	var err error
	_, err = client.Lock("test") //大概等待5秒
	log.Println(err)
	wait.Wait()
}

// 过期时间 WithExpiry 默认8秒
func TestClient_LockWithExpiry(t *testing.T) {
	mutex, err := client.Lock("test", redsync.WithExpiry(time.Second))
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second * 2)
	unlock, err := mutex.Unlock()
	fmt.Println(unlock, err) //锁已过期 返回 false nil
}

// 设置取锁次数 WithTries 默认32次 WithRetryDelay 每次取锁的等待时间
func TestClient_LockWithTries(t *testing.T) {
	go func() {
		mutex, err := client.Lock("test")
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 6)
		unlock, err := mutex.Unlock()
		fmt.Println(unlock, err) //锁已过期 返回 false nil
	}()
	time.Sleep(time.Second)
	_, err := client.Lock("test", redsync.WithTries(2), redsync.WithRetryDelay(time.Second*2))
	fmt.Println(err)
}

func TestClient_Lock3(t *testing.T) {
	for j := 0; j < 1; j++ {
		count := 0
		wait := &sync.WaitGroup{}
		wait.Add(2000)
		for i := 0; i < 2000; i++ {
			go func(i int, wd *sync.WaitGroup) {
				mutex, err := client.Lock("test", redsync.WithTries(1000), redsync.WithTimeoutFactor(20))
				if err != nil {
					panic(err)
				}
				count += i
				mutex.Unlock()
				wait.Done()
			}(i, wait)
		}
		wait.Wait()
		fmt.Println(count)
	}
}
