package config

import (
	"fmt"

	"github.com/spf13/viper"

	"ginrequest/redisdb"
)

const CONFIG_FILE_NAME = "config"

type YamlConfig struct {
	App         APP       `yaml:"app"`
	RedisConfig RDBConfig `yaml:"redis"`
}

type APP struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type RDBConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	PoolSize int    `yaml:"poolsize"`
}

var (
	AppYamlConfig YamlConfig
)

func InitConfig() {

	// 记录到文件。
	// logFile, err := os.Create("gin.log")

	// if err != nil {
	// 	fmt.Println("log file create fail:", err)

	// }
	// gin.DefaultWriter = io.MultiWriter(logFile)

	//定义路由日志的格式
	// gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	// 	log.Printf("interface %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	// }

	readAppConfig()

	fmt.Println("---InitConfig---")
	err := redisdb.InitRedisClient(AppYamlConfig.RedisConfig.Host,
		AppYamlConfig.RedisConfig.Password,
		AppYamlConfig.RedisConfig.DB,
		AppYamlConfig.RedisConfig.PoolSize)

	if err != nil {
		fmt.Println("redis client error:", err)
	} else {
		fmt.Println("redis client success!")
	}
}

func writeAppConfig() {
	viperConfig := viper.New()
	// 设置配置文件名，没有后缀
	viperConfig.SetConfigFile("config.yaml")

	//viper.Set("yaml", "this is a example of yaml")
	/*
			app:
		  name: ksnowlv关于Gin框架学习实践
		  version: v1.0
		redis:
		  host: localhost:6379
		  password: ""
		  db: 0
		  poolsize: 100
	*/

	viperConfig.Set("app.name", "ksnowlv关于Gin框架学习实践")
	viperConfig.Set("app.version", "v1.0")

	viperConfig.Set("redis.host", "localhost:6379")
	viperConfig.Set("redis.password", "")
	viperConfig.Set("redis.db", 0)
	viperConfig.Set("redis.poolsize", 100)

	if err := viperConfig.WriteConfig(); err != nil {
		fmt.Println(err)
	}
}

func readAppConfig() {

	viperConfig := viper.New()
	// 设置配置文件名，没有后缀
	viperConfig.SetConfigFile("./config/config1.yaml")

	// 读取解析
	if err := viperConfig.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("配置文件未找到！%v\n", err)
			return
		} else {
			fmt.Printf("找到配置文件,但是解析错误,%v\n", err)
			return
		}
	}

	if err := viperConfig.Unmarshal(&AppYamlConfig); err != nil {
		fmt.Printf("配置映射错误,%v\n", err)
	}
	fmt.Printf("App: %+v\n", AppYamlConfig.App)
	fmt.Printf("RedisConfig: %+v\n", AppYamlConfig.RedisConfig)
}
