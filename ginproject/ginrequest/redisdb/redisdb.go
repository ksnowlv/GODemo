package redisdb

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var (
	RDB *redis.Client
)

// 初始化连接
func InitRedisClient(addr string, password string, db int, poolSize int) (err error) {

	RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
		PoolSize: poolSize, // 连接池大小
	})

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = RDB.Ping().Result()
	return err
}

func RedisGetString(key string) string {
	val, err := RDB.Get(key).Result()
	if err != nil {
		fmt.Printf("get value failed, err:%v\n", err)
	}

	return val
}

func demo() {

	err := RDB.Set("name", "ksnowlv", 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := RDB.Get("name").Result()
	if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	}
	fmt.Println("name:", val)

}
