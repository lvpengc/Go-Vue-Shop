package orm

// Category 类目表
type Category struct {
	Base
	Name    string `gorm:"column:name;type:varchar(63);not null;" json:"name" form:"name"`               // 名称
	Desc    string `gorm:"column:desc;type:varchar(255);not null;" json:"desc" form:"desc"`              // 类目简述
	IconURL string `gorm:"column:icon_url;type:varchar(255);not null;" json:"icon_url" form:"icon_url"`  // 类目图标
	PicURL  string `gorm:"column:pic_url;type:varchar(255);not null;" json:"pic_url" form:"pic_url"`     // 类目图片
	Sort    uint   `gorm:"column:sort;not null;default:100;" json:"sort" form:"sort"`                    // 排序值
	Status  uint8  `gorm:"column:status;type:tinyint(2);not null;default:1" json:"status" form:"status"` // 状态：1-正常，2-停用
}
