package orm

// Goods 商品表
type Goods struct {
	Base
	Name        string       `gorm:"column:name;type:varchar(63);not null;" json:"name" form:"name"`                                        // 商品名称，限定不超过10字
	GoodsSn     string       `gorm:"column:goods_sn;type:varchar(63);not null;" json:"goods_sn" form:"goods_sn"`                            // 商品编号
	CategoryID  uint64       `gorm:"column:category_id;not null;default:0;" json:"category_id" form:"category_id"`                          // 类目ID
	Gallery     string       `gorm:"column:gallery;not null;type:varchar(1023);" json:"gallery" form:"gallery"`                             // 商品宣传图片列表，采用JSON数组格式
	Keywords    string       `gorm:"column:keywords;not null;type:varchar(255);" json:"keywords" form:"keywords"`                           // 商品搜索关键字，逗号分隔
	Unit        string       `gorm:"column:unit;type:varchar(15);default:'件';" json:"unit" form:"unit"`                                     // 商品单位，例如件、盒
	Brief       string       `gorm:"column:brief;type:varchar(255);" json:"brief" form:"brief"`                                             // 商品简介
	Detail      string       `gorm:"column:detail;type:text;" json:"detail" form:"detail"`                                                  // 商品详细介绍，图文并茂，是富文本格式
	PicURL      string       `gorm:"column:detail;type:varchar(255);" json:"pic_url" form:"pic_url"`                                        // 商品页面商品图片
	ShareURL    string       `gorm:"column:detail;type:varchar(255);" json:"share_url" form:"share_url"`                                    // 商品分享朋友圈图片
	IsOnSale    uint8        `gorm:"column:is_on_sale;default:2;not null;" json:"is_on_sale" form:"is_on_sale"`                             // 是否上架：1-是，2-否
	IsNew       uint8        `gorm:"column:is_new;default:2;not null;" json:"is_new" form:"is_new"`                                         // 是否新品：1-是，2-否
	IsHot       uint8        `gorm:"column:is_hot;default:2;not null;" json:"is_hot" form:"is_hot"`                                         // 是否热卖：1-是，2-否
	Sort        uint         `gorm:"column:sort;not null;default:100;" json:"sort" form:"sort"`                                             // 排序值
	Price       float64      `gorm:"column:price;not null;type:decimal(10,2);default:0.00;" json:"price" form:"price"`                      // 红字价格，未选中规格前显示
	MarkedPrice float64      `gorm:"column:marked_price;not null;type:decimal(10,2);default:0.00;" json:"marked_price" form:"marked_price"` // 划线价
	GoodsAttrs  []GoodsAttr  `gorm:"save_associations:false;"`                                                                              // 商品拥有多种固定属性
	GoodsStocks []GoodsStock `gorm:"save_associations:false;"`                                                                              // 商品对于每种规格组合都有对应库存量
}
