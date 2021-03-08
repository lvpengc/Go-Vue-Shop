package orm

// OrderGoods 订单商品表
type OrderGoods struct {
	Base
	OrderID        uint64  `gorm:"column:order_id;not null;default:0;" json:"order_id" form:"order_id"`                            // 订单ID
	GoodsID        uint64  `gorm:"column:goods_id;not null;default:0;" json:"goods_id" form:"goods_id"`                            // 商品ID
	GoodsStockID   uint64  `gorm:"column:goods_stock_id;not null;default:0;" json:"goods_stock_id" form:"goods_stock_id"`          // 商品货品表的货品ID
	GoodsSn        string  `gorm:"column:goods_sn;type:varchar(63);not null;" json:"goods_sn" form:"goods_sn"`                     // 商品编号
	GoodsName      string  `gorm:"column:goods_name;type:varchar(127);not null;" json:"goods_name" form:"goods_name"`              // 商品名称
	Number         uint64  `gorm:"column:number;not null;default:0;" json:"number" form:"number"`                                  // 购买数量
	Price          float64 `gorm:"column:price;type:decimal(10,2);not null;default:0.00;" json:"price" form:"price"`               // 商品货品的售价
	Specifications string  `gorm:"column:specifications;type:varchar(1023);not null;" json:"specifications" form:"specifications"` // 商品规格value列表，采用JSON数组格式
	PicURL         string  `gorm:"column:pic_url;type:varchar(255);not null;" json:"pic_url" form:"pic_url"`                       // 商品图片，具体库存图片
	Status         uint8   `gorm:"column:status;type:tinyint(2);not null;default:1" json:"status" form:"status"`                   // 状态：1-待评价，2-已评价
}
