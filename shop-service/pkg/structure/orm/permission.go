package orm

// Permission 权限表
type Permission struct {
	Base
	MenuID uint64  `gorm:"column:menu_id;not null;default:0" json:"menu_id" form:"menu_id"`       // 菜单ID
	Name   string  `gorm:"column:name;type:varchar(255);not null;" json:"name" form:"name"`       // 权限资源名称，如 order
	Action string  `gorm:"column:action;type:varchar(255);not null;" json:"action" form:"action"` // 权限操作名称，如 read
	Roles  []Role  `gorm:"many2many:role_permission;save_associations:false;"`                    // 管理员和角色的多对多关联
	Menu   Menu    `gorm:"save_associations:false;"`                                              // 权限多对一归属于具体的页面菜单
	Routes []Route `gorm:"save_associations:false;"`                                              // 权限一对多包含具体的后端路由
}
