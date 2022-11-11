package cache

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type Redis struct {
	*redis.Client
}

func New(config *Config) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password, // no password set
		DB:       config.Db,       // use default DB
		PoolSize: config.PoolSize,
	})

	redisClient := &Redis{client}

	_, err := redisClient.Ping().Result()
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to redis")
	}

	return redisClient, nil
}

func (c *Redis) AddJwtToBlacklist(token string) error {
	_, err := c.Set(token, true, 0).Result()
	if err != nil {
		log.Fatalf("Could not connect to redis %v", err)
		return err
	}
	return nil
}

func (c *Redis) IsExists(token string) (bool, error) {
	rs, err := c.Exists(token).Result()
	if err != nil {
		log.Fatalf("Could not connect to redis %v", err)
		return rs > 0, err
	}
	return rs > 0, nil
}

func (c *Redis) SetValue(key string, value string, expiry time.Duration) error {
	err := c.Set(key, value, expiry).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Redis) SetValueAsInterface(key string, i interface{}, expiry time.Duration) error {
	err := c.Set(key, i, expiry).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Redis) GetValue(key string) (string, error) {
	value, err := c.Get(key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *Redis) GetValueAsInterface(key string, i interface{}) error {
	err := c.Get(key).Scan(i)
	if err != nil {
		return err
	}
	return nil
}

func (c *Redis) RemoveKey(key string) error {
	_, err := c.Del(key).Result()
	if err != nil {
		return err
	}
	return nil
}

func (c *Redis) GetTTL(key string) (time.Duration, error) {
	return c.TTL(key).Result()
}

func (c *Redis) IsTTLLessThen(key string, d time.Duration) (bool, error) {
	ttl, err := c.GetTTL(key)
	if err != nil {
		return false, err
	}
	if err != nil {
		return false, err
	}
	return ttl < d, nil
}
