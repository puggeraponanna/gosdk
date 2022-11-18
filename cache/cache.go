package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v9"
)

type SimpleCache interface {
	Get(ctx context.Context, key string, out any) error
	Set(ctx context.Context, key string, value any, expiry time.Duration) error
	SetOrReplace(ctx context.Context, key string, value any, expiry time.Duration) error
}

type SimpleRedisCache struct {
	client *redis.Client
}

func NewSimpleRedisCache(url string) SimpleCache {
	opts, err := redis.ParseURL(url)
	if err != nil {
		log.Fatal(err)
	}
	return &SimpleRedisCache{
		client: redis.NewClient(opts),
	}
}

func (src *SimpleRedisCache) Get(ctx context.Context, key string, out any) error {
	res, err := src.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(res), out)
	if err != nil {
		return err
	}
	return nil
}

func (src *SimpleRedisCache) Set(ctx context.Context, key string, value any, expiry time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = src.client.SetNX(ctx, key, v, expiry).Result()
	if err != nil {
		return err
	}
	return nil
}

func (src *SimpleRedisCache) SetOrReplace(ctx context.Context, key string, value any, expiry time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = src.client.Set(ctx, key, v, expiry).Result()
	if err != nil {
		return err
	}
	return nil
}
