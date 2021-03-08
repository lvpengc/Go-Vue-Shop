package orm

// GoodsAttr 商品参数表
type GoodsAttr struct {
	Base
	GoodsID             uint64               `gorm:"column:goods_id;not null;default:0;" json:"goods_id" form:"goods_id"` // 商品ID
	Attr                string               `gorm:"column:attr;type:varchar(255);" json:"attr" form:"attr"`              // 商品参数名称
	Value               string               `gorm:"column:value;type:varchar(255);" json:"value" form:"value"`           // 商品参数值，固定属性才有此值
	Type                uint8                `gorm:"column:type;not null;default:1;" json:"type" form:"type"`             // 参数类型：1-固定属性、2-可选规格
	GoodsSpecifications []GoodsSpecification `gorm:"save_associations:false;"`                                            // 对于可选规格可拥有多个可选值
}
