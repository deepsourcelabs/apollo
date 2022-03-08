package redis

import "github.com/go-redis/redis/v8"

type DBConfig struct {
	Addr string
}

func (c *DBConfig) InitDB() *redis.Client {
	redisOpts := &redis.Options{
		Addr:     c.Addr,
		Password: "",
		DB:       0,
	}
	db := redis.NewClient(redisOpts)
	return db
}
