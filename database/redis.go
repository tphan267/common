package database

import (
	"fmt"

	"github.com/tphan267/common/system"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

func InitRedis() {
	redisAddr := fmt.Sprintf("%s:6379", system.Env("REDIS_HOST", "localhost"))

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,                      // Redis server address
		Password: system.Env("REDIS_PASS", ""),   // No password set
		DB:       system.EnvAsInt("REDIS_DB", 1), // use default DB
	})
}
