# Redis Notes
- data storage/querying tool that stores key-object pairs in RAM to achieve very fast performance 
- typically used in addition to a database to improve lookup times via caching
  - especially for large, slow running queries or for data that doesn't change that often

## Contents 
- [Redis Notes](#redis-notes)
  - [Contents](#contents)
  - [installation](#installation)
  - [How to use it](#how-to-use-it)

## installation 
- Raw dog it on WSL:
> wsl --install (in powershell)
>
> wsl -d ubuntu 
>
>$ sudo apt install redis
>
>$ redis-cli
- or if using with Go:
> go mod init github.com/myrepo

> go get github.com/redis/go-redis/v9
- (must be done after initializing a module)
- then in code:

            import "github.com/redis/go-redis/v9"

## How to use it
- connect to a new client 

            rdb := redis.NewClient(&redis.Options{
                Addr: "localhost:6379",
                Password: "",
                DB: 0, //0 is default
                Protocol: 2,
            })

            ctx := context.Background()

- OR connect using a connection string:

            opt, err := redis.ParseURL("redis://<user>:<password>@localhost:6379/<db>")

            if err != nil {
                panic(err)
            }

            client := redis.NewClient(opt)

- after connecting, test connection by storing and receiving a simple string:

            err := rdb.Set(ctx, "foo", "bar", 0).Err()

            if err != nil {
                panic(err)
            }

            val, err := rdb.Get(ctx, "foo").Result()

            if err != nil {
                panic(err)
            }

            fmt.Println("foo", val)

- example storing/retrieving a hash using `HSet` and `HGet`:

            hashFields := []string{
                "model", "deimos",
                "brand", "Ergonom",
                "type", "enduro bikes",
                "price", "4972",
            }

            res1, err := rdb.HSet(ctx, "bike:1", hashFields).Result()

            if err != nil {
                panic(err)
            }

            fmt.Println(res1) //output will be 4

            res2, err := rdb.HGet(ctx, )