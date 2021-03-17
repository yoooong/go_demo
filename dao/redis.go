package dao

import (
	"context"
	"douyu-new/dao/cache"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var RedisDao *cache.RedisRDB

func RedisInit() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//var cancel context.CancelFunc

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println("Redis Ping =", pong, err)
	if err != nil {
		panic("redis ping error")
	}

	RedisDao = NewRedisDao(rdb, ctx)
	fmt.Println(" RedisRDB", RedisDao)
}

func NewRedisDao(rdb *redis.Client, ctx context.Context) (RedisDao *cache.RedisRDB) {
	RedisDao = &cache.RedisRDB{
		RDB: rdb,
		Ctx: ctx,
	}
	return
}
