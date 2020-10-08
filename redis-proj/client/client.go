package client

import (
	"github.com/go-redis/redis"
)
type Client struct {
	Redis *redis.Client
}

func (c *Client) NewRedisClient () *redis.Client{

client :=	redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	return client
}
