package orm

// AdminRole 用户对应角色表
type AdminRole struct {
	Base
	AdminID uint64 `gorm:"column:admin_id;not null;default:0" json:"admin_id" form:"admin_id"` // 管理账号ID
	RoleID  uint64 `gorm:"column:role_id;not null;default:0" json:"role_id" form:"role_id"`    // 角色ID
}
