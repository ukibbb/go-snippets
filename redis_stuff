package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type Cacher interface {
	Get(int) (string, error)
	Remove(int) error
	Set(int, string) error
}

type NopCache struct{}

func (c NopCache) Get(int) (string, error) { return "", nil }
func (c NopCache) Remove(int) error        { return nil }
func (c NopCache) Set(int, string) error   { return nil }

type Store struct {
	cache Cacher
	data  map[int]string
}

func NewStore(c Cacher) *Store {
	return &Store{
		cache: c,
		data: map[int]string{
			1: "Donald Trump is new old president of US",
			2: "Foo is not bar and bar is not buzz",
			3: "Must watch AnthonyGG",
		},
	}
}

func (s *Store) Set() {}
func (s *Store) Get(key int) (string, error) {
	val, err := s.cache.Get(key)
	if err != nil {
		// busting the cache.
		if err := s.cache.Remove(key); err != nil {
			fmt.Println(err)
		}
		return val, nil
	}

	val, ok := s.data[key]
	if !ok {
		return "", fmt.Errorf("key not found %d", key)
	}

	if err := s.cache.Set(key, val); err != nil {
		return "", err
	}
	fmt.Println("returning key from internal storage")

	return val, nil
}

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(c *redis.Client) *RedisCache {
	return &RedisCache{
		client: c,
	}
}

func (c *RedisCache) Get(key int) (string, error) {
	ctx := context.Background()
	return c.client.Get(ctx, strconv.Itoa(key)).Result()

}
func (c RedisCache) Remove(key int) error {
	ctx := context.Background()
	_, err := c.client.Del(ctx, strconv.Itoa(key)).Result()
	return err

}
func (c *RedisCache) Set(key int, value string) error {
	ctx := context.Background()
	_, err := c.client.Set(ctx, strconv.Itoa(key), value, 0).Result()
	return err

}
