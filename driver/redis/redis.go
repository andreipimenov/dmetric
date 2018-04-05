package redis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	Pool *redis.Pool
}

func NewRedis(server string, db int) (*Redis, error) {
	r := &Redis{
		Pool: getRedisPool(server, db),
	}
	if err := r.Ping(); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Redis) Close() error {
	return r.Pool.Close()
}

func getRedisPool(server string, db int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     10000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			_, err = c.Do("SELECT", db)
			if err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func (r *Redis) Ping() error {
	c := r.Pool.Get()
	defer c.Close()
	msg, err := redis.String(c.Do("PING"))
	if msg != "PONG" {
		return fmt.Errorf("server responsed %s on PING request", msg)
	}
	return err
}

func (r *Redis) Set(key string, value string) error {
	c := r.Pool.Get()
	_, err := c.Do("SET", key, value)
	c.Close()
	return err
}

func (r *Redis) Get(key string) (string, error) {
	c := r.Pool.Get()
	defer c.Close()
	return redis.String(c.Do("GET", key))
}
