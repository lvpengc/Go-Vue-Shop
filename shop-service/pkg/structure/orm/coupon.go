package orm

import "time"

// Coupon 优惠券信息及规则表
type Coupon struct {
	Base
	Name       string     `gorm:"column:name;type:varchar(63);not null;" json:"name" form:"name"`                                 // 优惠券名称，如限时满减券
	Desc       string     `gorm:"column:desc;type:varchar(255);not null;" json:"desc" form:"desc"`                                // 优惠券介绍，通常是显示优惠券使用限制文字
	Total      int64      `gorm:"column:total;not null;default:-1;" json:"total" form:"total"`                                    // 优惠券数量：-1-无限量
	UserLimit  int64      `gorm:"column:user_limit;not null;default:1;" json:"user_limit" form:"user_limit"`                      // 用户领券限制数量：-1-无限制，其他代表具体数量
	Amount     float64    `gorm:"column:amount;type:decimal(10,2);not null;default:0.00" json:"amount" form:"amount"`             // 优惠金额
	ConsumMin  float64    `gorm:"column:consum_min;type:decimal(10,2);not null;default:0.00" json:"consum_min" form:"consum_min"` // 最低消费金额
	Type       uint8      `gorm:"column:type;not null;default:1;" json:"type" form:"type"`                                        // 优惠券类型：1-通用券，用户领取；2-注册赠券；3-优惠券码兑换
	Status     uint8      `gorm:"column:status;type:tinyint(2);not null;default:1" json:"status" form:"status"`                   // 优惠券状态：1-正常，2-停用下架，3-过期失效
	GoodsType  uint8      `gorm:"column:goods_type;not null;default:1;" json:"goods_type" form:"goods_type"`                      // 商品限制类型，1-全场通用，2-类目限制，3-商品限制
	GoodsValue string     `gorm:"column:goods_value;not null;type:varchar(1023);" json:"goods_value" form:"goods_value"`          // 商品限制值，goods_type 所对应的ID限制，逗号分隔或json数组
	TimeType   uint8      `gorm:"column:time_type;not null;default:1" json:"time_type" form:"time_type"`                          // 时间限制类型，1-时间不限；2-基于领取时间的有效天数days；3-start_date和end_date是优惠券有效期；
	Days       uint       `gorm:"column:days;type:smallint(6);default:0;" json:"days" form:"days"`                                // 基于领取时间的有效天数
	StartDate  time.Time  `gorm:"column:start_date;type:datetime;" json:"start_date" form:"start_date"`                           // 使用券开始时间
	EndDate    time.Time  `gorm:"column:end_date;type:datetime;" json:"end_date" form:"end_date"`                                 // 使用券截至时间
	CouponCode CouponCode `gorm:"save_associations:false;"`                                                                       // 一种优惠券可以生成多个兑换码
}
