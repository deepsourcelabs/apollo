package redis

import (
	"github.com/burntcarrot/apollo/utils"
	"github.com/go-redis/redis/v8"
)

// GetConn returns a Redis connection with configurations.
func GetConn(config *utils.Config) *redis.Client {
	redisOpts := &redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	}

	db := redis.NewClient(redisOpts)
	return db
}
