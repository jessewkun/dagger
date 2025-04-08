package router

import (
	"dagger/app/demo"
	"dagger/app/user"
	"dagger/common"
	"dagger/lib/middleware"
	"dagger/lib/mysql"
	"dagger/lib/redis"
	"dagger/lib/sys"
	"dagger/service"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "dagger/docs"

	swaggerFiles "github.com/swaggo/files"
)

// 跨域配置
var crosConfig = middleware.CrosConfig{
	AllowedOrigins: map[string]bool{
		"http://localhost:8001": true, // 后端
		"http://localhost:5173": true, // 前端
		"http://127.0.0.1:8001": true, // 后端
		"http://127.0.0.1:5173": true, // 前端
	},
	AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "OPTIONS"},
	AllowHeaders: []string{"Content-Type, Authorization, Content-Length,Keep-Alive,credentials,Cache-Control,user,X-Requested-With,If-Modified-Since,Cache-Control,Pragma,Last-Modified,Accept,Accept-Encoding,Accept-Language,Connection,Host,Referer,User-Agent,Origin,Sec-Ch-Ua,Sec-Ch-Ua-Mobile,Sec-Ch-Ua-Platform,Sec-Fetch-Dest,Sec-Fetch-Mode,Sec-Fetch-Site"},
}

func InitRouter(r *gin.Engine) *gin.Engine {
	// 全局中间件
	r.Use(middleware.DaggerRecovery(), middleware.Cros(crosConfig), middleware.Trace(), middleware.IOLog())
	r.NoMethod(HandleNotFound)
	r.NoRoute(HandleNotFound)

	// 注册系统路由
	registerSystemRoutes(r)

	// 注册API路由
	registerAPIRoutes(r)

	// 注册静态资源路由
	registerStaticRoutes(r)

	return r
}

// registerSystemRoutes 注册系统相关路由
func registerSystemRoutes(r *gin.Engine) {
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
}

// registerAPIRoutes 注册API路由
func registerAPIRoutes(r *gin.Engine) {
	var checkLoginFunc middleware.CheckLoginFunc = service.NewUserService().DecodeToken
	// var needLoginFunc middleware.NeedLoginFunc = service.NewUserService().DecodeTokenWithNoError

	// 用户相关路由
	v1UserRouter := r.Group("/user/v1/")
	v1UserRouter.POST("/auth", user.AuthHandler)
	v1UserRouter.POST("/logout", user.LogoutHandler)
	v1UserRouter.GET("/checkLogin", middleware.CheckLogin(checkLoginFunc), user.CheckLoginHandler)
	v1UserRouter.POST("/modify", middleware.CheckLogin(checkLoginFunc), user.ModifyHandler)
	v1UserRouter.POST("/modify/password", middleware.CheckLogin(checkLoginFunc), user.ModifyPasswordHandler)
	v1UserRouter.POST("/reset/password", user.ResetPasswordHandler)
	v1UserRouter.POST("/set/password", user.SetPasswordHandler)

	// Demo相关路由
	v1Proxy := r.Group("/demo/v1/")
	v1Proxy.GET("/index", demo.IndexHandler)
	v1Proxy.GET("/one", demo.OneHandler)
	v1Proxy.GET("/list", demo.ListHandler)
	v1Proxy.POST("/add", demo.AddHandler)
	v1Proxy.POST("/update", demo.UpdateHandler)
	v1Proxy.POST("/delete", demo.DeleteHandler)
	v1Proxy.POST("/redis", demo.RedisHandler)
}

// registerStaticRoutes 注册静态资源路由
func registerStaticRoutes(r *gin.Engine) {
	// 用于访问静态文件，比如图片、视频等非前端页面静态文件
	r.Static("/resources", "./")

	// 用于访问前端页面静态文件
	if common.IsRelease() {
		// 生产环境使用 nginx 反向代理，静态资源由 nginx 提供
		r.LoadHTMLFiles("/var/www/static/dist/index.html")
	} else {
		r.Static("/static", "./static/dist")
		r.LoadHTMLFiles(".static/dist/index.html")
	}

	// 主页路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// 解决 vue router createWebHistory 路由刷新问题，该模式路由不带 # 号，直接访问非 / 路由会返回 404
	// 解决办法1 服务端配置，所有路由返回 index.html
	// 解决办法2 vue 项目配置 createWebHashHistory，路由带 # 号
	r.NoRoute(func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
}

func HandleNotFound(c *gin.Context) {
	c.Status(http.StatusNotFound)
}
