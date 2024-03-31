package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Conf struct {
	Address   string `yaml:"address" env-default:"0.0.0.0:6379"`
	Password  string `yaml:"password"`
	DefaultDb int    `yaml:"defaultDb" env-default:"0"`
}

type Context struct {
	client *redis.Client
}

func NewRedisContext(conf Conf) *Context {
	return &Context{
		client: redis.NewClient(&redis.Options{
			Addr:     conf.Address,
			Password: conf.Password,
			DB:       conf.DefaultDb,
		}),
	}
}

func (receiver *Context) Get(key string) *redis.StringCmd {
	return receiver.client.Get(context.Background(), key)
}

func (receiver *Context) Set(key string, value string, ttl time.Duration) *redis.StatusCmd {
	return receiver.client.Set(context.Background(), key, value, ttl)
}

func (receiver *Context) GetAllKeys() *redis.Cmd {
	return receiver.client.Do(context.Background(), "KEYS", "*")
}

func (receiver *Context) Subscribe(pattern string) *redis.PubSub {
	return receiver.client.PSubscribe(context.Background(), pattern)
}
