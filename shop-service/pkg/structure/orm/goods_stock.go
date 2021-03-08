package orm

// GoodsStock 商品货品表
type GoodsStock struct {
	Base
	GoodsID        uint64  `gorm:"column:goods_id;not null;default:0;" json:"goods_id" form:"goods_id"`                            // 商品ID
	Specifications string  `gorm:"column:specifications;type:varchar(1023);not null;" json:"specifications" form:"specifications"` // 商品规格value列表，采用JSON数组格式
	Price          float64 `gorm:"column:price;not null;type:decimal(10,2);default:0.00;" json:"price" form:"price"`               // 价格
	StockNum       uint64  `gorm:"column:stock_num;not null;default:0;" json:"stock_num" form:"stock_num"`                         // 库存
	PicURL         string  `gorm:"column:detail;type:varchar(255);" json:"pic_url" form:"pic_url"`                                 // 商品货品图片
}
