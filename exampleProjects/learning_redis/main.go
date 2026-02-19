// go mod init github.com/fake-module
// go get github.com/redis/go-redis/v9
// start a Redis server locally

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func Main() {
	fmt.Println("Hello Redis")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //typically not hardcoded
		Password: "",               //default password
		DB:       0,                //default DB
	})

	ctxFiveSeconds, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel() //don't forget to close the context to avoid memory leaks

	fmt.Println("Attempting to ping localhost Redis server...")
	ping, err := client.Ping(ctxFiveSeconds).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Ping Response:", ping)

	//now we want to set a key and then get it back
	err = client.Set(ctxFiveSeconds, "apples", "red", 0).Err() //0 means the key will not expire, otherwise it denotes expiration time in seconds
	if err != nil {
		panic(err)
	}
	fmt.Println("\nSet key 'apples' to 'red'.\nAccessing key 'apples'...")

	val, err := client.Get(ctxFiveSeconds, "apples").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("'apples' is: ", val)

	//saving more than one key at a time is done using structs, Hashes, or MSet and MGet

	//setting multiple keys at once using structs/JSON
	type Fruit struct {
		Name    string `redis:"name"`  //these are annotations of the form `user:"fieldname"` that tell clients how to map struct fields to other types of data
		Color   string `redis:"color"` //this would also work with `json:"fieldname"` for JSON encoding/decoding
		Barcode string `redis:"-"`     //annotations of "-" will be ignored by the Redis client and not stored in Redis, but the field name will still appear
	}
	ban := Fruit{Name: "banana", Color: "yellow"}
	jsonString, err := json.Marshal(ban)
	if err != nil {
		panic(err)
	}
	err = client.Set(ctxFiveSeconds, "banana", jsonString, 0).Err()
	if err != nil {
		panic(err)
	}

	//accessing the struct/JSON key returns a JSON string that we can unmarshal back into a struct
	val, err = client.Get(ctxFiveSeconds, "banana").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nResult of accessing the stored struct/JSON string: %s\n\n", val)

	//setting multiple keys at once using MSet
	err = client.MSet(ctxFiveSeconds, "grape", "purple", "cucumber", "green").Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("\nSet multiple keys at once using MSet.\nAccessing keys 'grape' and 'cucumber'...")

	//accessing multiple keys at once using MGet
	vals, err := client.MGet(ctxFiveSeconds, "grape", "cucumber").Result() //vals return type is []interface{}
	if err != nil {
		panic(err)
	}
	fmt.Println("'grape' and 'cucumber' are: ")
	for _, val := range vals {
		fmt.Println(val)
	}
	fmt.Print("\n")

	//setting multiple keys at once using a hash
	hashFields := []string{
		"name", "grape",
		"color", "purple",
		"barcode", "123456789",
	}

	res1, err := client.HSet(ctxFiveSeconds, "fruit:grape", hashFields).Result() //naming keys to hashes like "fruit:grape" is a common convention to reduce confusion: "type:key"
	if err != nil {
		panic(err)
	}
	//res1 is the number of fields that were added to the hash, so in this case it should be 3
	fmt.Println("Set multiple keys at once using a hash. Number of fields added to the hash:", res1)

	//you can access individual hash fields using HGet(context, keyToHashInRedis, fieldName)
	singleHashVal, err := client.HGet(ctxFiveSeconds, "fruit:grape", "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("The 'name' field in the hash is: ", singleHashVal)

	//or you can use HGetAll(context, hashKey) which returns a map[string]string of all the fields and values in the hash
	hashMap, err := client.HGetAll(ctxFiveSeconds, "fruit:grape").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("We received the hash: \n", hashMap)
}

func CacheIntoRedis(ctx context.Context, client *redis.Client, key string, value string, expiration time.Duration) error {
	err := client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
