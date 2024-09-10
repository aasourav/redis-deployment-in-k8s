package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

// Create a Redis client
var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.19.255.202:4899", // Redis server address
		Password: "aas",                 // No password set
	})

	// Ping Redis to check the connection
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Printf("Redis Ping: %s\n", pong)

	// Set a key-value pair in Redis
	err = rdb.Set(ctx, "mykey", "myvalue", 0).Err()
	if err != nil {
		log.Fatalf("Could not set key: %v", err)
	}

	// Get the value of the key from Redis
	val, err := rdb.Get(ctx, "mykey").Result()
	if err != nil {
		log.Fatalf("Could not get key: %v", err)
	}
	fmt.Printf("mykey: %s\n", val)

	// Test getting a non-existent key
	val2, err := rdb.Get(ctx, "nonexistent").Result()
	if err == redis.Nil {
		fmt.Println("Key does not exist")
	} else if err != nil {
		log.Fatalf("Error getting key: %v", err)
	} else {
		fmt.Printf("nonexistent: %s\n", val2)
	}
}
