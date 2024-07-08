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
	r.Use(gin.Recovery(), middleware.Cros(), middleware.Trace(), middleware.IOLog())
	r.NoMethod(HandleNotFound)
	r.NoRoute(HandleNotFound)

	// ping
	r.GET("/healthcheck/ping", func(c *gin.Context) {
		c.String(200, "Hello World")
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

	v1Proxy := r.Group("/demo/v1/").Use(middleware.CheckLogin())
	v1Proxy.POST("/index", demo.IndexHandler)

	return r
}

func HandleNotFound(c *gin.Context) {
	c.Status(http.StatusNotFound)
}
