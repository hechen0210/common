/**
@Time : 2019/11/25 15:33
@Author : hechen
@File : redis
@Software: GoLand
*/
package redis

import "github.com/go-redis/redis"

type Config struct {
	Host string
	Port string
	Auth string
}

type Redis struct {
	client *redis.Client
	Error  error
}

func (c Config) New() Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Host + ":" + c.Port,
		Password: c.Auth,
	})
	_, err := client.Ping().Result()
	return Redis{
		client: client,
		Error:  err,
	}
}
