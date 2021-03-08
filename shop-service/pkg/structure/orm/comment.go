package orm

// Comment 评论表
type Comment struct {
	Base
	UserID       uint64 `gorm:"column:user_id;not null;default:0;" json:"user_id" form:"user_id"`                      // 订单ID
	GoodsID      uint64 `gorm:"column:goods_id;not null;default:0;" json:"goods_id" form:"goods_id"`                   // 商品ID
	OrderGoodsID uint64 `gorm:"column:order_goods_id;not null;default:0;" json:"order_goods_id" form:"order_goods_id"` // 订单商品ID
	Content      string `gorm:"column:content;not null;type:varchar(1023);" json:"content" form:"content"`             // 评论文字内容
	PicURLs      string `gorm:"column:pic_urls;not null;type:varchar(1023);" json:"pic_urls" form:"pic_urls"`          // 图片地址列表，JSON数组
	Star         uint8  `gorm:"column:star;not null;default:5;" json:"star" form:"star"`                               // 星星评分，1-5
}
