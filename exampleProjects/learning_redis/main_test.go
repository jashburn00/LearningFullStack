package main

import (
	"context"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

var wg sync.WaitGroup

func TestMain(t *testing.T) {
	//spoofing
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	defer r.Close()
	defer w.Close()
	os.Stdout = w

	wg.Go(func() {
		t.Run("Main should run without panicking", func(t *testing.T) {
			Main()
		})
	})

	wg.Wait()
	//unspoofing
	os.Stdout = originalStdout
}

func TestCacheIntoRedis(t *testing.T) {
	//setup
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	wg.Go(func() {
		t.Run("CacheIntoRedis should cache a key-value pair without error", func(t *testing.T) {
			err := CacheIntoRedis(ctx, client, "url", "fakeurl.com/abcd", 0)
			if err != nil {
				t.Errorf("error encountered: %s", err.Error()) // error: context deadline exceeded
			}
		})
	})

	wg.Go(func() {
		t.Run("CacheIntoRedis should obey the given time to live duration", func(t *testing.T) {
			err := CacheIntoRedis(ctx, client, "Now You See Me", "John Cena", time.Second*1)
			if err != nil {
				t.Errorf("error encountered: %s", err.Error())
			}
			time.Sleep(time.Second * 2) //timing out a Redis key's TTL removes both the key and value
			val, err := client.Get(ctx, "Now You See Me").Result()
			if val != "" {
				t.Errorf("error encountered: val was %q but should be empty", val)
			}
		})
	})

	wg.Wait()
}
