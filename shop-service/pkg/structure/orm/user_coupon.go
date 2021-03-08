package orm

import "time"

// UserCoupon 优惠券用户使用表
type UserCoupon struct {
	Base
	UserID   uint64    `gorm:"column:user_id;not null;default:0;" json:"user_id" form:"user_id"`             // 订单ID
	CouponID uint64    `gorm:"column:coupon_id;not null;default:0;" json:"coupon_id" form:"coupon_id"`       // 优惠券ID
	OrderID  uint64    `gorm:"column:order_id;not null;default:0;" json:"order_id" form:"order_id"`          // 订单ID
	UsedAt   time.Time `gorm:"column:used_at;type:datetime;" json:"used_at" form:"used_at"`                  // 使用时间
	StartAt  time.Time `gorm:"column:start_at;type:datetime;" json:"start_at" form:"start_at"`               // 有效期开始时间
	EndAt    time.Time `gorm:"column:end_at;type:datetime;" json:"end_at" form:"end_at"`                     // 有效期截至时间
	Status   uint8     `gorm:"column:status;type:tinyint(2);not null;default:1" json:"status" form:"status"` // 状态：1-未使用， 2-已使用，3-过期失效
}
