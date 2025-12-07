package common

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

func NewCache(client *redis.Client) *Cache {
	return &Cache{client: client}
}

// SET CACHE
func (c *Cache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, bytes, ttl).Err()
}

//GET CACHE 
//
// Attempts to read cached value into "target" object.
// Returns:
//
//	(true, nil)  → cache HIT
//	(false, nil) → cache MISS
//	(false, err) → Redis/JSON error
func (c *Cache) Get(ctx context.Context, key string, target interface{}) (bool, error) {

	val, err := c.client.Get(ctx, key).Result()

	if err == redis.Nil {
		return false, nil 
	}

	if err != nil {
		return false, err
	}

	if len(val) == 0 {
		return false, nil
	}

	err = json.Unmarshal([]byte(val), target)
	if err != nil {
		return false, err
	}

	return true, nil 
}
