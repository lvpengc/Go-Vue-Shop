package router

import (
	"github.com/hellozouyou/easy-mall/admin_api/controller/auth"

	"github.com/gin-gonic/gin"
)

// RegisteRouterAuth 注册user登录相关的路由
func RegisteRouterAuth(app *gin.Engine) {
	g := app.Group("/auth")
	g.GET("/login", auth.Login)
	g.GET("/logout")
	g.GET("/info")
	g.GET("/refresh")
	g.GET("/update")
}
