package test

import (
	"dagger/common"
	"dagger/lib/logger"
	"dagger/lib/mysql"
	"flag"
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

var configFile string

func TestMain(m *testing.M) {
	flag.StringVar(&configFile, "c", "config.toml", "config file path")
	flag.Parse()

	configFile = "/Users/wangkun/Documents/localhost/go/src/dagger/config/debug.toml"
	viper.SetConfigFile(configFile)
	fmt.Println("Loading config file " + configFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("viper ReadInConfig %s error: %s\n", configFile, err))
	}

	if err := viper.Unmarshal(&common.DaggerCfg); err != nil {
		panic(fmt.Errorf("viper Unmarshal %s error: %s\n", configFile, err))
	}

	logger.InitLogger()
	mysql.InitDb()

	m.Run()
}
