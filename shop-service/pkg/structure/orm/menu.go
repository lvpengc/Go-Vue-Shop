package orm

// Menu 前端路由菜单表
type Menu struct {
	Base
	Path           string       `gorm:"column:path;type:varchar(255);" json:"path" form:"path"`                                  // 当前路由的path，只是这一级的
	Component      string       `gorm:"column:component;type:varchar(255);" json:"component" form:"component"`                   // 当前路由所依赖的组件
	Hidden         uint8        `gorm:"column:hidden;" json:"hidden" form:"hidden"`                                              // 该路由是否在侧边栏隐藏：1-是，2-否
	AlwaysShow     uint8        `gorm:"column:alwaysShow;" json:"alwaysShow" form:"alwaysShow"`                                  // 是否总是显示根路由：1-是，2-否
	Redirect       string       `gorm:"column:redirect;type:varchar(255);" json:"redirect" form:"redirect"`                      // 在面包屑中点击会重定向去的地址
	Name           string       `gorm:"column:name;type:varchar(255);" json:"name" form:"name"`                                  // 路由名称，必须添加
	Type           uint8        `gorm:"column:type;" json:"type" form:"type"`                                                    // 分类，用于多套导航菜单：1-管理平台，2-财务账单，3-短信业务
	MetaTitle      string       `gorm:"column:meta_title;type:varchar(255);" json:"meta_title" form:"meta_title"`                // 该路由在侧边栏和面包屑中展示的名称
	MetaIcon       string       `gorm:"column:meta_icon;type:varchar(255);" json:"meta_icon" form:"meta_icon"`                   // 该路由的图标
	MetaNocache    uint8        `gorm:"column:meta_noCache;" json:"meta_noCache" form:"meta_noCache"`                            // 该页面是否不使用<keep-alive> 缓存：1-是，2-否
	MetaBreadcrumb uint8        `gorm:"column:meta_breadcrumb;" json:"meta_breadcrumb" form:"meta_breadcrumb"`                   // 该页面是否在面包屑中展示：1-是，2-否
	MetaAffix      uint8        `gorm:"column:meta_affix;" json:"meta_affix" form:"meta_affix"`                                  // 该路由是否固定到标签导航上：1-是，2-否
	MetaRoles      string       `gorm:"column:meta_roles;type:varchar(255);" json:"meta_roles" form:"meta_roles"`                // 该路由允许查看的角色，优先于permission的判断，逗号分隔
	MetaActiveMenu string       `gorm:"column:meta_activeMenu;type:varchar(255);" json:"meta_activeMenu" form:"meta_activeMenu"` // 该路由在侧边栏active的实际菜单项
	ParentID       uint64       `gorm:"column:parent_id;type:varchar(255);" json:"parent_id" form:"parent_id"`                   // 上级路由的ID
	Sort           uint         `gorm:"column:sort;not null;default:0;" json:"sort" form:"sort"`                                 // 同级路由的排序值
	Permissions    []Permission `gorm:"save_associations:false;" `                                                               //  // 角色和权限的多对多关联
}
