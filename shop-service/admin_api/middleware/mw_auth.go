package middleware

import (
	"time"

	"github.com/micro/go-micro/util/log"

	"github.com/hellozouyou/easy-mall/pkg/constant"
	"github.com/hellozouyou/easy-mall/pkg/helper"
	"github.com/hellozouyou/easy-mall/pkg/jwt"
	"github.com/hellozouyou/easy-mall/pkg/redis"

	"github.com/gin-gonic/gin"
)

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			jwtToken jwt.Token
			err      error
		)
		// 判断是否传入了跳过判断的函数，如若有且满足则直接return
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}
		// 从Header头或获取token
		if token := c.GetHeader(constant.TOKENKEY); token != "" {
			// 解析token获取用户信息
			jwtToken, err = jwt.Parse(token)
			if err != nil {
				log.Errorf("token parse fail：%v, err: %v", token, err)
				helper.ResForbidden(c, (helper.InvalidTokenResponse()))
				return
			}
			if !time.Unix(jwtToken.Exp, 0).After(time.Now()) {
				log.Errorf("token expired：%v, %v, err: %v", token, jwtToken, err)
				helper.ResForbidden(c, helper.ExpiredTokenResponse())
				return
			}
		}

		if jwtToken.UUID != "" {
			//查询用户ID
			uid, err := redis.Client.GetUint64(jwtToken.UUID)
			if err != nil {
				log.Error("uuid get from cache fail：", uid, err)
				helper.ResForbidden(c, helper.InvalidTokenResponse())
				return
			}
			c.Set(constant.USERUUIDKEY, jwtToken.UUID)
			c.Set(constant.USERIDKEY, uid)
		}
		helper.ResUnauthorized(c, helper.InvalidTokenResponse())
		return
	}
}
