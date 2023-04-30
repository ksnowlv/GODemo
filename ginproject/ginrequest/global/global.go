package global

import (
	"database/sql"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

type XAppConfig struct {
	App    XAPP    `mapstructure:"app"`
	Logger XLogger `mapstructure:"log"`
	Redis  XRedis  `mapstructure:"redis"`
	MySQL  XMySQL  `mapstructure:"mysql"`
}

type XAPP struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Build   int    `yaml:"build"`
}

type XRedis struct {
	Host     string        `mapstructure:"host"`
	Password string        `mapstructure:"password"`
	DB       int           `mapstructure:"db"`
	PoolSize int           `mapstructure:"poolsize"`
	Timeout  time.Duration `mapstructure:"timeout"`
}

type XMySQL struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
}

// log:
//   level: "debug" # 日志级别 这里设置 debug 级别
//   filename: "web_app.log"
//   maxsize: 500 # 日志最大容量
//   maxage: 30 # 备份存储最大时间
//   maxbackups: 10 # 备份最大数量

type XLogger struct {
	Mode       string `mapstructure:"mode"`
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxsize"`
	MaxAge     int    `mapstructure:"maxage"`
	MaxBackups int    `mapstructure:"maxbackups"`
}

var (
	GRedis     *redis.Client
	GViper     *viper.Viper
	GAppConfig *XAppConfig
	GMySQL     *sql.DB
)
