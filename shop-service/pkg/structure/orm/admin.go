package orm

import "time"

// Admin 管理员表
type Admin struct {
	Base
	Name     string    `gorm:"column:name;type:varchar(63);not null;" json:"name" form:"name"`               // 登录账号
	Email    string    `gorm:"column:email;type:varchar(63);not null;" json:"email" form:"email"`            // 邮箱名
	Mobile   string    `gorm:"column:mobile;type:char(11);not null;" json:"mobile" form:"mobile"`            // 手机号码
	Nickname string    `gorm:"column:nickname;type:varchar(63);not null;" json:"nickname" form:"nickname"`   // 用户昵称
	Password string    `gorm:"column:password;type:varchar(255);not null;" json:"password" form:"password"`  // 密码
	Avatar   string    `gorm:"column:avatar;type:varchar(255);not null;" json:"avatar" form:"avatar"`        // 头像图片地址
	LoginAt  time.Time `gorm:"column:login_at;type:datetime;" json:"login_at" form:"login_at"`               // 登录时间
	Status   uint8     `gorm:"column:status;type:tinyint(2);not null;default:1" json:"status" form:"status"` // 状态：1-正常， 2-停用，3-软删
	Roles    []Role    `gorm:"many2many:admin_role;save_associations:false;"`                                // 管理员和角色的多对多关联
}
