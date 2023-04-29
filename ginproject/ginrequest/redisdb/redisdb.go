package redisdb

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"

	"ginrequest/global"
)

// 初始化redis连接
func InitRedisClient() (err error) {

	rc := &global.GAppConfig.Redis
	global.GRedis = redis.NewClient(&redis.Options{
		Addr:     rc.Host,
		Password: rc.Password, //密码
		DB:       rc.DB,       //默认db
		PoolSize: rc.PoolSize, // 连接池大小
	})

	timeoutCTX, cancelFunc := context.WithTimeout(context.Background(), rc.Timeout)
	defer cancelFunc()

	_, err = global.GRedis.Ping(timeoutCTX).Result()
	return err
}

func RedisSetString(key string, value string, t int64) bool {
	expire := time.Duration(t) * time.Second
	err := global.GRedis.Set(context.Background(), key, value, expire).Err()
	if err != nil {
		fmt.Printf("RedisSetString:%v", err)
		return false
	}
	return true
}

func RedisGetString(key string) string {

	v, err := global.GRedis.Get(context.Background(), key).Result()
	if err != nil {
		fmt.Printf("get value failed, err:%v\n", err)
		return ""
	}

	return v
}

// MSet 批量设置key的值
func RedisSetMultiValues(values ...interface{}) {
	err := global.GRedis.MSet(context.Background(), values).Err()

	if err != nil {
		fmt.Printf("--- RedisSetMultiValues--%v", err)
	}
}

// MGet 批量查询key的值
func RedisGetMultiValues(k ...string) []interface{} {
	res, err := global.GRedis.MGet(context.Background(), k...).Result()

	if err != nil {
		fmt.Printf("--- RedisGetMultiValues--%v", err)
	}

	return res
}
