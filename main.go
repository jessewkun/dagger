package main

import (
	"dagger/common"
	"dagger/lib/logger"
	"dagger/lib/mysql"
	"dagger/router"
	"flag"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "c", "config.yml", "config file path")
	flag.Parse()

	viper.SetConfigFile(configFile)
	fmt.Println("Loading config file " + configFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("viper ReadInConfig %s error: %s\n", configFile, err))
	}
	if err := viper.Unmarshal(&common.Cfg); err != nil {
		panic(fmt.Errorf("viper Unmarshal %s error: %s\n", configFile, err))
	}
	// 监控配置文件是否变化
	// viper 会自动监控配置文件的变化，当配置文件发生变化时，viper 会自动更新配置信息
	// 但是 viper 不会自动更新结构体，所以需要手动更新结构体
	// 这里只是为了 debug config 可以被动态更新，其他情况下不建议使用
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
	})

	logger.InitLogger(common.Cfg.Log)
	mysql.InitMysql(common.Cfg.Mysql)
	// redis.InitRedis(common.Cfg.Redis)
}

func main() {
	gin.SetMode(common.Cfg.Mode)
	r := gin.Default()

	router.InitRouter(r)
	r.Run(common.Cfg.Port)
}
