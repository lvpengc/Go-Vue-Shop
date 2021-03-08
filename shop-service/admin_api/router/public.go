package router

import "github.com/gin-gonic/gin"

// RegisteRouterPublic 注册可公开访问的路由
func RegisteRouterPublic(app *gin.Engine) {
	g := app.Group("/common")
	g.GET("/test")
}
