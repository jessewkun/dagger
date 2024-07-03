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

	if err := viper.Unmarshal(&common.DaggerCfg); err != nil {
		panic(fmt.Errorf("viper Unmarshal %s error: %s\n", configFile, err))
	}
	// 监控配置文件是否变化
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
	})
	viper.WatchConfig()

	logger.InitLogger()
	mysql.InitDb()
}

func main() {
	gin.SetMode(common.DaggerCfg.Mode)
	r := gin.Default()

	router.Setup(r)
	r.Run(common.DaggerCfg.Port)
}
