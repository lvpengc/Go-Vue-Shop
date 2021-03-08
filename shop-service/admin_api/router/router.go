// Package router 定义路由列表
package router

import (
	//"html/template"

	"time"

	"github.com/hellozouyou/easy-mall/admin_api/middleware"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

// Init 路由初始化
func Init(app *gin.Engine) {

	// 解决跨域问题
	app.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, POST, PATCH, OPTIONS, DELETE, PUT, HEAD, TRACE, CONNECT",
		RequestHeaders:  "Origin, Authorization, Content-Type, x-token",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	// 全局传入需要校验auth的中间件，并定义不需要校验登录的url
	notCheckLoginURLArr := []string{
		"/auth/login",
		"/auth/logout",
	}
	app.Use(middleware.UserAuthMiddleware(
		middleware.AllowPathPrefixSkipper(notCheckLoginURLArr...),
	))

	// 注册路由
	RegisteRouterAuth(app)
	RegisteRouterPublic(app)
}
