package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var a = 1

// Login 登录
func Login(c *gin.Context) {
	fmt.Println(a)
	a++
}
