// Package orm 声明数据库的表结构体
package orm

import "time"

// Base orm基础模型
type Base struct {
	ID        uint64    `gorm:"column:id;primary_key;auto_increment;" json:"id" form:"id"`                     // 主键
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;" json:"created_at" form:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null;" json:"updated_at" form:"updated_at"` // 更新时间
}
