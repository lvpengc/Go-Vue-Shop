package orm

// Route 后端路由表
type Route struct {
	Base
	Path         string     `gorm:"column:path;type:varchar(255);" json:"path" form:"path"`                            // 后端路由，如：/admin/:id
	PermissionID uint64     `gorm:"column:permission_id;not null;default:0" json:"permission_id" form:"permission_id"` // 权限ID
	Permission   Permission `gorm:"save_associations:false;"`                                                          // 一组路由操作归属于某一个具体权限
}
