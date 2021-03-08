package orm

// Collect 商品收藏表
type Collect struct {
	Base
	UserID  uint64 `gorm:"column:user_id;not null;default:0;" json:"user_id" form:"user_id"`    // 订单ID
	GoodsID uint64 `gorm:"column:goods_id;not null;default:0;" json:"goods_id" form:"goods_id"` // 商品ID
}
