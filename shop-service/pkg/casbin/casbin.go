// Package casbin casbin.go 定义casbin的权限操作
package casbin

import (
	"strconv"

	"github.com/hellozouyou/easy-mall/pkg/mysql"
	"github.com/hellozouyou/easy-mall/pkg/structure/orm"

	"github.com/casbin/casbin"
	"github.com/micro/go-micro/util/log"
)

// 前缀定义
const (
	prefixUserID = "u"
	prefixRoleID = "r"
)

// Enforcer Enforcer
var (
	Enforcer *casbin.Enforcer
)

// InitEnforcer 角色-URL导入
func InitEnforcer() (err error) {
	if Enforcer, err = casbin.NewEnforcerSafe(casbin.NewModel(casbinModel)); err != nil {
		log.Errorf("InitEnforcer NewEnforcerSafe err: %v", err)
		return
	}
	var roles []orm.Role
	if err = mysql.Client.Find(&roles).Error; err != nil {
		log.Errorf("InitEnforcer get roles err: %v", err)
		return
	}
	for _, role := range roles {
		setRolePermission(Enforcer, role.ID)
	}
	return
}

// setRolePermission 设置角色权限，查询角色对应权限所属下的具体路由，实际判断路由是否匹配
func setRolePermission(enforcer *casbin.Enforcer, roleid uint64) {
	// mysql.Client.Find()
	// var rolemenus []tablemodels.RoleMenu
	// err := models.Find(&tablemodels.RoleMenu{RoleID: roleid}, &rolemenus)
	// if err != nil {
	// 	return
	// }
	// for _, rolemenu := range rolemenus {
	// 	menu := tablemodels.Menu{}
	// 	where := tablemodels.Menu{}
	// 	where.ID = rolemenu.PermissionID
	// 	_, err = models.First(&where, &menu)
	// 	if err != nil {
	// 		return
	// 	}
	// 	if menu.MenuType == 3 {
	// 		// 获取配置
	// 		enforcer.AddPermissionForUser(prefixRoleID+strconv.Itoa(int(roleid)), menu.URL,
	// 			"GET|POST|PUT")
	// 	}
	// }
}

// DeleteRole 删除角色
func DeleteRole(roleids []uint64) {
	if Enforcer == nil {
		return
	}
	for _, rid := range roleids {
		Enforcer.DeletePermissionsForUser(prefixRoleID + strconv.Itoa(int(rid)))
		Enforcer.DeleteRole(prefixRoleID + strconv.Itoa(int(rid)))
	}
}

// CsbinSetRolePermission 设置角色权限
func CsbinSetRolePermission(roleid uint64) {
	if Enforcer == nil {
		return
	}
	Enforcer.DeletePermissionsForUser(prefixRoleID + strconv.Itoa(int(roleid)))
	setRolePermission(Enforcer, roleid)
}

// CsbinCheckPermission 检查用户是否有权限
func CsbinCheckPermission(userID, url, methodtype string) (bool, error) {
	return Enforcer.EnforceSafe(prefixUserID+userID, url, methodtype)
}

// CsbinAddRoleForUser 用户角色处理
func CsbinAddRoleForUser(userid uint64) (err error) {
	// if Enforcer == nil {
	// 	return
	// }
	// uid := prefixUserID + strconv.Itoa(int(userid))
	// Enforcer.DeleteRolesForUser(uid)
	// var roles []orm.Roles
	// err = models.Find(&tablemodels.AdminsRole{AdminID: userid}, &roles)
	// if err != nil {
	// 	return
	// }
	// for _, adminsrole := range adminsroles {
	// 	Enforcer.AddRoleForUser(uid, prefixRoleID+strconv.Itoa(adminsrole.RoleID))
	// }
	return
}
