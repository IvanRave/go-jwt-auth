package main

import (
	"fmt"
	"time"
	"gopkg.in/redis.v4"
)

// Client is a Redis client representing a pool of zero or more underlying connections. It's safe for concurrent use by multiple goroutines.
var pool *redis.Client

func initPool() error {
	// NewClient returns a client to the Redis Server
	pool = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 10, // default
	})

	pong, err := pool.Ping().Result()

	fmt.Println(pong)
	
	return err
}

func checkKey() error {
	err := pool.Set("mykey", "myvalue", 0).Err()

	if err != nil {
		return err
	}

	val, err := pool.Get("mykey").Result()
	if err != nil {
		return err
	}

	fmt.Println("key", val)
	
	return nil
}

// setLoginAndCode with expiration time
// Returns:
// true - if set (not exists)
// false - if not set (exists already)
func setLoginAndCode(lgn string, code string) (bool, error) {
	// NX = not exists
	// Zero expiration means the key has no expiration time.
	return pool.SetNX(lgn, code, 10*time.Second).Result()
}

// getVcode by email/phone
func getVcode(lgn string) (string, error) {

	// Get returns *StringCmd
	// contains filtered or unexported fields	
	str, err := pool.Get(lgn).Result()

	if err != nil{
		if err == redis.Nil {
			// return empty string
			return "", nil
		}

		return "", err
	}
	
	return str, nil
}
