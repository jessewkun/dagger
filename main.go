package main

import (
	"context"
	"dagger/common"
	"dagger/lib/logger"
	"dagger/lib/mysql"
	"dagger/router"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	srv := &http.Server{
		Addr:    common.Cfg.Port,
		Handler: router.InitRouter(r),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Errorf("Server startup failed: %v", err))
		}
	}()

	gracefulExit(srv)
}

// gracefulExit 优雅退出
func gracefulExit(srv *http.Server) {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-exit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	logger.Info(ctx, "main", "Received signal: %v. Shutting down server...", sig)

	if err := srv.Shutdown(ctx); err != nil {
		logger.ErrorWithMsg(ctx, "main", "Server shutdown failed: %v", err)
	}
	logger.Info(ctx, "main", "Server gracefully shutdown")
	fmt.Println("Server gracefully shutdown")
}
