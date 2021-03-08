// Package helper web.go api接口的相关辅助工具
package helper

import (
	"net/http"

	"github.com/hellozouyou/easy-mall/pkg/constant"

	"github.com/gin-gonic/gin"
)

// APIResponse api 接口的响应结构
type APIResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 预定义不变常用的api响应量
var (
	invalidTokenResponse = &APIResponse{
		Code: constant.InvalidToken,
		Msg:  "Invalid Token",
	}
	expiredTokenResponse = &APIResponse{
		Code: constant.InvalidToken,
		Msg:  "Invalid Token",
	}
)

// InvalidTokenResponse 预定义无效token的响应
func InvalidTokenResponse() *APIResponse {
	return invalidTokenResponse
}

// ExpiredTokenResponse 预定义无效token的响应
func ExpiredTokenResponse() *APIResponse {
	return expiredTokenResponse
}

// ResJSON 返回json格式
func ResJSON(c *gin.Context, code int, v *APIResponse) {
	c.JSON(code, v)
	c.Abort()
}

// ResSucc 返回成功
func ResSucc(c *gin.Context) {
	res := &APIResponse{
		Code: constant.Succ,
		Msg:  "Success",
		Data: struct{}{},
	}
	ResJSON(c, http.StatusOK, res)
}

// ResSuccWithData 返回成功，携带数据
func ResSuccWithData(c *gin.Context, v interface{}) {
	res := &APIResponse{
		Code: constant.Succ,
		Msg:  "Success",
		Data: v,
	}
	ResJSON(c, http.StatusOK, res)
}

// ResSuccWithMsg 返回成功、自定义消息文案
func ResSuccWithMsg(c *gin.Context, msg string) {
	res := &APIResponse{
		Code: constant.Succ,
		Msg:  msg,
		Data: struct{}{},
	}
	ResJSON(c, http.StatusOK, res)
}

// ResSuccWithDataMsg 返回成功、自定义消息文案、数据
func ResSuccWithDataMsg(c *gin.Context, msg string, v interface{}) {
	res := &APIResponse{
		Code: constant.Succ,
		Msg:  msg,
		Data: v,
	}
	ResJSON(c, http.StatusOK, res)
}

// ResForbidden 返回 403 Forbidden
func ResForbidden(c *gin.Context, res *APIResponse) {
	ResJSON(c, http.StatusForbidden, res)
}

// ResUnauthorized 返回 401 Unauthorized
func ResUnauthorized(c *gin.Context, res *APIResponse) {
	ResJSON(c, http.StatusUnauthorized, res)
}
