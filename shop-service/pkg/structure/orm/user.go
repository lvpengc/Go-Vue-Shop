package orm

import "time"

// User 用户表
type User struct {
	Base
	OpenID        string      `gorm:"column:openid;type:varchar(63);not null;" json:"openid" form:"openid"`                 // 微信登录openid
	SessionKey    string      `gorm:"column:session_key;type:varchar(255);not null;" json:"session_key" form:"session_key"` // 微信登录会话KEY
	Gender        uint8       `gorm:"column:gender;type:tinyint(2);not null;default:1" json:"gender" form:"gender"`         // 性别：1-未知，2-男，3-女
	Birthday      time.Time   `gorm:"column:birthday;type:date" json:"birthday" form:"birthday"`                            // 生日
	Nickname      string      `gorm:"column:nickname;type:varchar(63);not null;" json:"nickname" form:"nickname"`           // 用户昵称
	Phone         string      `gorm:"column:password;type:char(11);not null;" json:"phone" form:"phone"`                    // 用户手机号码
	Avatar        string      `gorm:"column:avatar;type:varchar(255);not null;" json:"avatar" form:"avatar"`                // 用户头像图片地址
	LoginAt       time.Time   `gorm:"column:login_at;type:datetime;" json:"login_at" form:"login_at"`                       // 最近登录时间
	Status        uint8       `gorm:"column:status;type:tinyint(2);not null;default:1" json:"status" form:"status"`         // 状态：1-正常，2-停用，3-软删
	UserAddresses UserAddress `gorm:"save_associations:false;"`
	UserCoupons   UserCoupon  `gorm:"save_associations:false;"`
}
