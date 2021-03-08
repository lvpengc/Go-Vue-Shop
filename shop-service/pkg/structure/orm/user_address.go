package orm

// UserAddress 用户收货地址
type UserAddress struct {
	Base
	UserID   uint64 `gorm:"column:user_id;not null;default:0;" json:"user_id" form:"user_id"`             // 订单ID
	Name     string `gorm:"column:name;type:varchar(63);not null;" json:"name" form:"name"`               // 收货人
	Mobile   string `gorm:"column:mobile;type:varchar(20);not null;" json:"mobile" form:"mobile"`         // 电话号码，可以固话
	Province string `gorm:"column:province;type:varchar(63);not null;" json:"province" form:"province"`   // 行政区域表的省ID
	City     string `gorm:"column:city;type:varchar(63);not null;" json:"city" form:"city"`               // 行政区域表的市ID
	County   string `gorm:"column:county;type:varchar(63);not null;" json:"county" form:"county"`         // 行政区域表的区县ID
	Detail   string `gorm:"column:detail;type:varchar(127);not null;" json:"detail" form:"detail"`        // 详细收货地址
	AreaCode uint64 `gorm:"column:area_code;" json:"area_code" form:"area_code"`                          // 地区代号
	PostCode string `gorm:"column:post_code;type:char(10);" json:"post_code" form:"post_code"`            // 邮编
	Status   uint8  `gorm:"column:status;type:tinyint(2);not null;default:1" json:"status" form:"status"` // 状态：1-非默认， 2-默认
}
