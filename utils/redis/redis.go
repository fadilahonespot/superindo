package redis

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheService interface {
	Set(ctx context.Context, key, value string, duration time.Duration) (err error) 
	Get(ctx context.Context, key string) (value string, err error)
	Delete(ctx context.Context, key string) (err error)
}

type cache struct {
	client *redis.Client
}

func NewCache() CacheService {
	fmt.Println("Connect Redis Client.....")

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")), // Redis server address
		Password: os.Getenv("REDIS_PASSWORD"),                                            // No password for your local setup
		Username: os.Getenv("REDIS_USERNAME"),                                            // Username
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	return &cache{client: client}
}

func (w *cache) Set(ctx context.Context, key, value string, duration time.Duration) (err error) {
	log.Printf("[CACHED SET] key: %v, value: %v \n", key, value)
	err = w.client.Set(ctx, key, value, duration).Err()
	return
}

func (w *cache) Get(ctx context.Context, key string) (value string, err error) {
	log.Printf("[CACHED GET] key: %v \n", key)
	err = w.client.Get(ctx, key).Scan(&value)
	return
}

func (w *cache) Delete(ctx context.Context, key string) (err error) {
	log.Printf("[CACHED DEL] key: %v \n", key)
	err = w.client.Del(ctx, key).Err()
	return
}
