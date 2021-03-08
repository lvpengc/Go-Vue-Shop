package orm

// ProvinceCity 省市表
type ProvinceCity struct {
	Base
	Name       string `gorm:"column:name;type:varchar(16);" json:"name" form:"name"`                     // 省/市名称
	ParentID   string `gorm:"column:parent_id;type:varchar(8);" json:"parent_id" form:"parent_id"`       // 父级ID
	SimpleName string `gorm:"column:simple_name;type:varchar(8);" json:"simple_name" form:"simple_name"` // 简称
	OpType     uint8  `gorm:"column:op_type;type:tinyint(2);" json:"op_type" form:"op_type"`             // 类型：1-省，2-市
	AreaCode   uint64 `gorm:"column:area_code;" json:"area_code" form:"area_code"`                       // 地区代号
	PostCode   string `gorm:"column:post_code;type:char(10);" json:"post_code" form:"post_code"`         // 邮编
	PinYin     string `gorm:"column:pinyin;type:char(30);" json:"pinyin" form:"pinyin"`                  // 拼音
}
