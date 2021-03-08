package orm

import "time"

// Order 订单表
type Order struct {
	Base
	UserID       uint64       `gorm:"column:user_id;not null;default:0;" json:"user_id" form:"user_id"`                            // 用户ID
	OrderSn      string       `gorm:"column:order_sn;type:varchar(63);not null;" json:"order_sn" form:"order_sn"`                  // 订单编号
	Consignee    string       `gorm:"column:consignee;type:varchar(63);not null;" json:"consignee" form:"consignee"`               // 收货人名称
	Mobile       string       `gorm:"column:mobile;type:varchar(63);not null;" json:"mobile" form:"mobile"`                        // 收货人手机号
	Address      string       `gorm:"column:address;type:varchar(255);not null;" json:"address" form:"address"`                    // 收货具体地址
	Message      string       `gorm:"column:message;type:varchar(1023);not null;" json:"message" form:"message"`                   // 用户订单留言
	GoodsMoney   float64      `gorm:"column:goods_money;type:decimal(10,2);not null;" json:"goods_money" form:"goods_money"`       // 商品总费用
	FreightMoney float64      `gorm:"column:freight_money;type:decimal(10,2);not null;" json:"freight_money" form:"freight_money"` // 配送费用
	CouponMoney  float64      `gorm:"column:coupon_money;type:decimal(10,2);not null;" json:"coupon_money" form:"coupon_money"`    // 优惠券减免
	OrderMoney   float64      `gorm:"column:order_money;type:decimal(10,2);not null;" json:"order_money" form:"order_money"`       // 订单费用， goods_money + freight_money - coupon_money
	PayID        string       `gorm:"column:pay_id;type:varchar(63);" json:"pay_id" form:"pay_id"`                                 // 微信付款编号
	PayAt        time.Time    `gorm:"column:pay_at;type:datetime;" json:"pay_at" form:"pay_at"`                                    // 微信付款时间
	ConfirmAt    time.Time    `gorm:"column:confirm_at;type:datetime;" json:"confirm_at" form:"confirm_at"`                        // 用户确认收货时间
	EndAt        time.Time    `gorm:"column:end_at;type:datetime;" json:"end_at" form:"end_at"`                                    // 订单关闭时间
	ShipAt       time.Time    `gorm:"column:ship_at;type:datetime;" json:"ship_at" form:"ship_at"`                                 // 发货开始时间
	ShipSn       string       `gorm:"column:ship_sn;type:varchar(63);not null;" json:"ship_sn" form:"ship_sn"`                     // 发货编号
	ShipChannel  string       `gorm:"column:ship_channel;type:varchar(63);not null;" json:"ship_channel" form:"ship_channel"`      // 发货快递公司
	Status       uint8        `gorm:"column:status;type:tinyint(2);not null;default:101" json:"status" form:"status"`              // 状态：101-未付款；102-未付款，自行取消；103-未付款，系统取消；201-已支付，待发货；202-已支付，取消订单，申请退款；203-已支付，退款完成；301-已发货，运输中；401-已收货，用户自行确认，402-已收货，系统超时自动确认, 501-已评价
	OrderGoodses []OrderGoods `gorm:"save_associations:false;"`
}
