// Package constant constant.go 定义项目常量
package constant

import "time"

// TimeZone 全局时区
var TimeZone, _ = time.LoadLocation("Asia/Shanghai")

// DEFAULTEXPIRE 默认有效期，秒
const DEFAULTEXPIRE int64 = 3 * 24 * 60 * 60 // 3天

// 服务名
const (
	WEBADMIN    string = "easymall.web.admin"
	WEBADMINAPI string = "easymall.web.admin_api"
	WEBAPPAPI   string = "easymall.web.app_api"
)

// MQ 消息的 topic
const ()

// 业务常量
const (
	HASHSALT            = "github.com/hellozouyou/easy-mall"
	SUPERADMINID uint64 = 1000
	TOKENKEY            = "Bearer-Token" //页面token键名
	USERIDKEY           = "X-USERID"     //页面用户ID键名
	USERUUIDKEY         = "X-UUID"       //页面UUID键名
)
