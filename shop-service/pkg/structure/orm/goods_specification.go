package orm

// GoodsSpecification 商品规格表
type GoodsSpecification struct {
	Base
	GoodsID     uint64 `gorm:"column:goods_id;not null;default:0;" json:"goods_id" form:"goods_id"`                // 商品ID
	GoodsAttrID uint64 `gorm:"column:goods_attr_id;not null;default:0;" json:"goods_attr_id" form:"goods_attr_id"` // 商品属性ID
	Value       string `gorm:"column:value;type:varchar(255);" json:"value" form:"value"`                          // 商品规格值
}
