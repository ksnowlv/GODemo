package redisdb

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"

	"ginrequest/global"
)

// 初始化redis连接
func InitRedisClient() (err error) {

	rc := global.GAppYamlConfig.RedisConfig
	// fmt.Printf("host:%s,password:%s,db:%d,poolsize:%d",
	// 	rc.Host,
	// 	rc.Password,
	// 	rc.DB,
	// 	rc.PoolSize)
	global.GRedis = redis.NewClient(&redis.Options{
		Addr:     rc.Host,
		Password: rc.Password, // no password set
		DB:       rc.DB,       // use default DB
		PoolSize: rc.PoolSize, // 连接池大小
	})

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = global.GRedis.Ping().Result()
	return err
}

func RedisSetString(key string, value string, t int64) bool {
	expire := time.Duration(t) * time.Second
	err := global.GRedis.Set(key, value, expire).Err()
	if err != nil {
		fmt.Println("RedisSetString %v", err)
		return false
	}
	return true
}

func RedisGetString(key string) string {
	value, err := global.GRedis.Get(key).Result()
	if err != nil {
		fmt.Printf("get value failed, err:%v\n", err)
		return ""
	}

	return value
}

func demo() {

	err := global.GRedis.Set("name", "ksnowlv", 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := global.GRedis.Get("name").Result()
	if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	}
	fmt.Println("name:", val)

}
