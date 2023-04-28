package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

type YamlConfig struct {
	App         XAPP         `mapstructure:"app"`
	RedisConfig XRedisConfig `mapstructure:"redis"`
}

type XAPP struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Build   int    `yaml:"build"`
}

type XRedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"poolsize"`
}

var (
	GRedis         *redis.Client
	GViper         *viper.Viper
	GAppYamlConfig *YamlConfig
)
