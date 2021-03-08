package orm

// CouponCode 优惠券的兑换码表
type CouponCode struct {
	Base
	CouponID uint64 `gorm:"column:coupon_id;not null;"`                                                   // 兑换码优惠券ID
	UserID   uint64 `gorm:"column:user_id;not null;default:0;" json:"user_id" form:"user_id"`             // 用户ID，兑换后绑定到用户ID
	Code     string `gorm:"column:code;type:varchar(63);not null;" json:"code" form:"code"`               // 优惠券兑换码
	Status   uint8  `gorm:"column:status;type:tinyint(2);not null;default:1" json:"status" form:"status"` // 优惠券状态：1-正常，2-停用，3-已兑换
}
