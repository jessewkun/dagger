package router

import (
	"dagger/app/demo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) *gin.Engine {
	r.Use(gin.Recovery(), Cros())
	r.NoMethod(HandleNotFound)
	r.NoRoute(HandleNotFound)

	v1Proxy := r.Group("/demo/v1/").Use(CheckSign())
	v1Proxy.POST("/index", demo.IndexHandler)

	return r
}

func HandleNotFound(c *gin.Context) {
	c.Redirect(http.StatusFound, "https://www.baidu.com")
}
