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
	// TODO 部分配置变化后重置对应的配置
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
