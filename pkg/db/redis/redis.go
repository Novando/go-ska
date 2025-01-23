package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/novando/go-ska/pkg/logger"
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis struct {
	rdb    *redis.Client
	logger *logger.Logger
}

type Config struct {
	Port int
	Host string
	Pass string
	Log  *logger.Logger
}

// Init initiate Redis library
func Init(cfg *Config) *Redis {
	log := logger.Call()
	if cfg.Log != nil {
		log = cfg.Log
	}
	if cfg.Port == 0 {
		cfg.Port = 6379
	}
	if cfg.Host == "" {
		cfg.Host = "localhost"
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", cfg.Host, cfg.Port),
		Password: cfg.Pass,
		DB:       0, // use default DB
	})

	status := rdb.Ping(context.Background())
	if status.Err() != nil {
		log.Fatalf("Error connecting to redis: %s", status.Err())
	}

	log.Infof("Redis client initialized successfully.")

	return &Redis{
		rdb,
		log,
	}
}

// FlushAll drop all Redis data
func (r *Redis) FlushAll() {
	r.rdb.FlushAll(context.Background())
}

// Get Redis data value by the key, return error
// and empty string if key-value not exist
func (r *Redis) Get(key string) (val string, err error) {
	val, err = r.rdb.Get(context.Background(), key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			err = nil
		} else {
			r.logger.Errorf(fmt.Sprintf("Error getting value from redis: %s", err))
		}
	}
	return
}

// Set assign/create Redis value to it's key along with the expiration time.
// return error and empty string if key-value not exist
func (r *Redis) Set(key string, value string, expiration time.Duration) error {
	_, err := r.rdb.Set(context.Background(), key, value, expiration).Result()
	if err != nil {
		return fmt.Errorf("%s: %s", "Error setting value in redis", err)
	}
	return nil
}

// GetHash get Redis data value by the key on the certain hash, return error
// and empty string if key-value not exist
func (r *Redis) GetHash(key string, field string) (string, error) {
	val, err := r.rdb.HGet(context.Background(), key, field).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}
		return "", fmt.Errorf("%s: %s", "Error getting value from redis", err)
	}
	return val, nil
}

// SetHash assign/create Redis value to it's key on certain hash.
// return error and empty string if key-value not exist
func (r *Redis) SetHash(key string, field string, value string, ttl ...time.Duration) {
	c := context.Background()
	if _, err := r.rdb.HSet(c, key, field, value).Result(); err != nil {
		r.logger.Errorf(fmt.Sprintf("Error setting value in redis: %s", err))
	} else if len(ttl) > 0 {
		if err = r.rdb.Expire(c, field, ttl[0]).Err(); err != nil {
			r.logger.Errorf(fmt.Sprintf("Error setting TTL on hash: %s", err))
		}
		if err = r.rdb.HExpire(c, key, ttl[0], field).Err(); err != nil {
			r.logger.Errorf(fmt.Sprintf("Error setting TTL on field: %s", err))
		}
	}
	return
}

// Delete force delete Redis data value of non-hash key,
// Return error if no key were found.
func (r *Redis) Delete(key string) error {
	_, err := r.rdb.Del(context.Background(), key).Result()
	if err != nil {
		return fmt.Errorf("%s: %s", "Error deleting value from redis", err)
	}
	return nil
}

// Close Redis connection
func (r *Redis) Close() {
	defer func() {
		err := r.rdb.Close()
		if err != nil {
			r.logger.Fatalf("%s: %s", "Error closing redis", err)
		}
	}()
}
