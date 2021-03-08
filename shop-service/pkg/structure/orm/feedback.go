package orm

// Feedback 意见反馈表
type Feedback struct {
	Base
	UserID   uint64 `gorm:"column:user_id;not null;default:0;" json:"user_id" form:"user_id"`             // 用户ID，兑换后绑定到用户ID
	UserName string `gorm:"column:username;type:varchar(63);not null;" json:"username" form:"username"`   // 用户名称
	Mobile   string `gorm:"column:mobile;type:varchar(20);not null;" json:"mobile" form:"mobile"`         // 手机号，可以固话
	Content  string `gorm:"column:content;type:varchar(1023);not null;" json:"content" form:"content"`    // 具体内容
	PicURLs  string `gorm:"column:pic_urls;not null;type:varchar(1023);" json:"pic_urls" form:"pic_urls"` // 图片地址列表，JSON数组
}
