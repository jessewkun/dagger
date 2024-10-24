package router

import (
	"dagger/app/demo"
	"dagger/lib/middleware"
	"dagger/lib/mysql"
	"dagger/lib/redis"
	"dagger/lib/sys"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "dagger/docs"

	swaggerFiles "github.com/swaggo/files"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.DaggerRecovery(), middleware.Cros(), middleware.Trace(), middleware.IOLog())
	r.NoMethod(HandleNotFound)
	r.NoRoute(HandleNotFound)

	// ping
	r.GET("/healthcheck/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// 组件探活
	r.GET("/healthcheck/active", func(c *gin.Context) {
		data := map[string]interface{}{
			"db":    mysql.HealthCheck(),
			"cache": redis.HealthCheck(),
		}
		c.JSON(http.StatusOK, sys.SuccResp(c, data))
	})

	// swagger
	if gin.Mode() == gin.DebugMode {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	v1Proxy := r.Group("/demo/v1/")
	v1Proxy.GET("/index", demo.IndexHandler)
	v1Proxy.GET("/one", demo.OneHandler)
	v1Proxy.GET("/list", demo.ListHandler)
	v1Proxy.POST("/add", demo.AddHandler)
	v1Proxy.POST("/update", demo.UpdateHandler)
	v1Proxy.POST("/delete", demo.DeleteHandler)
	v1Proxy.POST("/redis", demo.RedisHandler)

	// 用于访问静态文件，比如图片、视频等非前端页面静态文件
	r.Static("/resources", "./")

	// 用于访问前端页面静态文件，目录替换为前端打包后的静态文件目录
	r.Static("/static", "./xplan-static/dist")

	// 使用go代理前端页面时候的主页，目录替换为前端打包后的静态文件目录
	r.LoadHTMLFiles("./xplan-static/dist/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// 解决 vue router createWebHistory 路由刷新问题，该模式路由不带 # 号，直接访问非 / 路由会返回 404
	// 解决办法1 服务端配置，所有路由返回 index.html
	// 解决办法2 vue 项目配置 createWebHashHistory，路由带 # 号
	r.NoRoute(func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	return r
}

func HandleNotFound(c *gin.Context) {
	c.Status(http.StatusNotFound)
}
