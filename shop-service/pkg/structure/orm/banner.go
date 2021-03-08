package orm

// Banner 首页 banner 表
type Banner struct {
	Base
	Name    string `gorm:"column:name;type:varchar(63);not null;" json:"name" form:"name"`               // 名称标题
	Link    string `gorm:"column:link;type:varchar(255);not null;" json:"link" form:"link"`              // 链接页面地址
	ImgURL  string `gorm:"column:img_url;type:varchar(255);not null;" json:"img_url" form:"img_url"`     // 图片地址
	Content string `gorm:"column:content;type:varchar(255);not null;" json:"content" form:"content"`     // 内容描述
	Sort    uint   `gorm:"column:sort;not null;default:100;" json:"sort" form:"sort"`                    // 排序值
	Status  uint8  `gorm:"column:status;type:tinyint(2);not null;default:1" json:"status" form:"status"` // 状态：1-正常，2-停用
}
