package services

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"

	var_api_gateway "github.com/abdullokh-mukhammadjonov/template_api_gateway/modules/template_variables/var_api_gateway"
)

type RedisManager interface {
	// The method for particular set caching
	Set(key string, value interface{}) error
	// The method for setting key-value pair with deadline
	SetWithDeadline(key string, value interface{}, duration time.Duration) error
	// The method for particular get from the cacher
	Get(key string, response interface{}) error
	//
	// The method to delete from the cache
	Del(key string) error
}

type redisClient struct {
	redis        *redis.Client
	services     ServiceManager
	timeDuration time.Duration
}

func NewRedisClient(client *redis.Client, services ServiceManager, duration time.Duration) RedisManager {
	return &redisClient{
		redis:        client,
		services:     services,
		timeDuration: duration,
	}
}

func (rc redisClient) SetEntityDrafts(response interface{}, filter []map[string]string) error {
	return nil
}

func (rc redisClient) GetEntityDrafts(response interface{}, filter []map[string]string) error {

	return nil
}
func (rc redisClient) Set(key string, value interface{}) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = rc.redis.Set(key, b, rc.timeDuration).Result()
	return err
}
func (rc redisClient) SetWithDeadline(key string, value interface{}, timeDuration time.Duration) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = rc.redis.Set(key, b, timeDuration).Result()
	return err
}
func (rc redisClient) Get(key string, response interface{}) error {
	b, err := rc.redis.Get(key).Bytes()
	if err == redis.Nil {
		return var_api_gateway.ErrCacheMiss
	} else if err != nil {
		return err
	}
	if b != nil {
		if err := json.Unmarshal(b, response); err != nil {
			return err
		}
	}
	return nil
}
func (rc redisClient) Del(key string) error {
	deleted, err := rc.redis.Del(key).Result()
	if err != nil {
		return err
	}
	if deleted == 0 {
		return var_api_gateway.ErrCacheMiss
	}
	return nil
}
