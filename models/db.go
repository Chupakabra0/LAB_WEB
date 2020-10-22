package models

import (
	"context"
	"github.com/go-redis/redis"
)

var (
	redisClient *redis.Client
	ctx = context.TODO()
)

func InitializeDB() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func GetData()  {
	
}