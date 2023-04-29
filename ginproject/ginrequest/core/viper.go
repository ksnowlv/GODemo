package core

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"ginrequest/global"
)

const CONFIG_YAML_FILE_NAME = "./config/config.yaml"

func InitViper() {
	// 实例化viper对象，指定文件路径以及对应的文件类型
	viperConfig := viper.New()
	viperConfig.SetConfigFile(CONFIG_YAML_FILE_NAME)
	viperConfig.SetConfigType("yaml")

	// 读配置文件
	err := viperConfig.ReadInConfig()
	if err != nil {
		fmt.Errorf("发生错误：%s \n", err)
	}

	// 监听配置文件的变化
	viperConfig.WatchConfig()
	// 每次修改配置文件中的配置项，都重新unmarshal一次
	viperConfig.OnConfigChange(func(evt fsnotify.Event) {
		fmt.Println("配置文件发生改变", evt.Name)
		//注意使用 & 引用
		updateConfigData(viperConfig)
	})

	updateConfigData(viperConfig)
}

func updateConfigData(v *viper.Viper) {
	err := v.Unmarshal(&global.GAppConfig)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("---config---%v\n", v)
		fmt.Printf("App: %+v\n", global.GAppConfig.App)
		fmt.Printf("log: %+v\n", global.GAppConfig.Logger)
		fmt.Printf("Redis: %+v\n", global.GAppConfig.Redis)
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
