package main

import (
	"dagger/common"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
)

var configfile string
var module string // 要调用的模块

func init() {
	flag.StringVar(&configfile, "c", "config.yml", "config file path")
	flag.StringVar(&module, "m", "module", "gin mode")
	flag.Parse()
}

func main() {
	if err := configor.Load(&common.DaggerCfg, configfile); err != nil {
		panic(err)
	}

	gin.SetMode(common.DaggerCfg.Mode)
	// ctx := context.Background()
	switch module {
	case "test":
		fmt.Printf("%+v\n", "test")
	default:
		fmt.Printf("%+v\n", "wrong module")
	}
}
