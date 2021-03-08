// Package middleware 定义各种中间件函数
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// NoMethodHandler 未找到请求方法的处理函数
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Method Not Allowed"})
	}
}

// NoRouteHandler 未找到请求路由的处理函数
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.HTML(http.StatusOK, "index.html", nil)
		c.JSON(http.StatusNotFound, gin.H{"message": "Route Not Found"})
	}
}

// SkipperFunc 定义中间件跳过函数
type SkipperFunc func(*gin.Context) bool

// AllowPathPrefixSkipper 检查请求路径是否包含指定的前缀
func AllowPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		// 判断当前url如若满足事先定义的任一url前缀，则返回true
		for _, p := range prefixes {
			if strings.HasPrefix(c.Request.URL.Path, p) {
				return true
			}
		}
		return false
	}
}
