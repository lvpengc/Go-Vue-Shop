package middleware

import (
	"net/http"

	"github.com/hellozouyou/easy-mall/pkg/casbin"
	"github.com/hellozouyou/easy-mall/pkg/constant"
	"github.com/hellozouyou/easy-mall/pkg/helper"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
)

// CasbinMiddleware casbin中间件
func CasbinMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 判断是否传入了跳过判断的函数，如若有且满足则直接return
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}
		// 获取auth中间件添加的uid参数
		uid, isExit := c.Get(constant.USERIDKEY)
		if !isExit {
			helper.ResForbidden(c, helper.InvalidTokenResponse())
			return
		}
		// 对超级管理员直接跳过
		if uid.(uint64) == constant.SUPERADMINID {
			c.Next()
			return
		}
		// 对当前路由和方法进行权限判断
		p := c.Request.URL.Path
		m := c.Request.Method
		if ok, err := casbin.CsbinCheckPermission(uid.(string), p, m); err != nil || !ok {
			log.Errorf("CsbinCheckPermission err: %v, ok: %v", err, ok)
			helper.ResJSON(c, http.StatusOK, &helper.APIResponse{
				Code: constant.NOPERMISSION,
				Msg:  "暂无访问权限",
			})
		}
		c.Next()
	}
}
