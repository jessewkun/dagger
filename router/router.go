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

	return r
}

func HandleNotFound(c *gin.Context) {
	c.Status(http.StatusNotFound)
}
