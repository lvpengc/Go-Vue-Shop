package orm

// Role 角色表
type Role struct {
	Base
	Name        string       `gorm:"column:name;type:varchar(255);not null;" json:"name" form:"name"` // 角色中文简称
	Remark      string       `gorm:"column:remark;type:varchar(255);" json:"remark" form:"remark"`    // 角色备注
	Admins      []Admin      `gorm:"many2many:admin_role;save_associations:false;"`                   // 管理员和角色的多对多关联
	Permissions []Permission `gorm:"many2many:role_permission;save_associations:false;"`              //  // 角色和权限的多对多关联
}
