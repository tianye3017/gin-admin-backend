package db

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/tianye3017/gin-admin-backend/config"
)

var Redis *redis.Client

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.SysConfig.Redis.Addr,
		Password: config.SysConfig.Redis.Password, // no password set
		DB:       config.SysConfig.Redis.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	} else {
		fmt.Println(pong, err)
		Redis = client
	}
}
