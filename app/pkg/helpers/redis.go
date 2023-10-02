package helpers

import (
	// "fmt"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func NewRedisClient() *redis.Client {
	return RedisClient
}

func InitRedisConn() (err error, client *redis.Client) {
	url := GetEnvStr("REDIS_URL")
	port := GetEnvStr("REDIS_PORT")
	pwd := GetEnvStr("REDIS_PASSWORD")

	RedisClient := redis.NewClient(&redis.Options{
		Addr:     url + ":" + port,
		Password: pwd, // no password set
		DB:       0,   // use default DB
	})

	// _, err = RedisClient.Ping().Result()
	// if err !=nil{
	// 	panic(err)
	// }
	return err, RedisClient
}
