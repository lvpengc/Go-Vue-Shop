package orm

// RolePermission 角色对应权限表
type RolePermission struct {
	Base
	RoleID       uint64 `gorm:"column:role_id;not null;default:0" json:"role_id" form:"role_id"`                   // 角色ID
	PermissionID uint64 `gorm:"column:permission_id;not null;default:0" json:"permission_id" form:"permission_id"` // 权限ID
}
