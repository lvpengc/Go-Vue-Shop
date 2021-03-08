package orm

// Issue 常见问题表
type Issue struct {
	Base
	Question string `gorm:"column:question;type:varchar(255);" json:"question" form:"question"` // 问题标题
	Answer   string `gorm:"column:answer;type:varchar(255);" json:"answer" form:"answer"`       // 问题答案
}
