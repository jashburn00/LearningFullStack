package main

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

// var wg sync.WaitGroup

func TestMain(t *testing.T) {
	//spoofing
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	defer r.Close()
	defer w.Close()
	os.Stdout = w

	t.Run("Main should run without panicking", func(t *testing.T) {
		//setup
		t.Parallel()
		//testing
		Main()
	})
	//unspoofing
	os.Stdout = originalStdout
}

func TestCacheIntoRedis(t *testing.T) {
	//setup

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	t.Run("CacheIntoRedis should cache a key-value pair without error", func(t *testing.T) {
		//setup
		t.Parallel()
		//testing
		err := CacheIntoRedis(context.Background(), client, "url", "fakeurl.com/abcd", 0)
		//verification
		if err != nil {
			t.Errorf("error encountered: %s", err.Error())
		}
	})
	t.Run("CacheIntoRedis should obey the given time to live duration", func(t *testing.T) {
		//setup
		t.Parallel()
		//testing
		err := CacheIntoRedis(context.Background(), client, "Now You See Me", "John Cena", time.Second*1)
		if err != nil {
			t.Errorf("error encountered: %s", err.Error())
		}
		time.Sleep(time.Second * 2) //timing out a Redis key's TTL removes both the key and value
		val, err := client.Get(context.Background(), "Now You See Me").Result()
		//verification
		if val != "" {
			t.Errorf("error encountered: val was %q but should be empty", val)
		}
	})
}
