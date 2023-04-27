/*
 * @Author: Profigogogogo wf851128@gmail.com
 * @Date: 2023-04-26 16:33:24
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2023-04-27 11:08:56
 * @FilePath: /go-sys/init/initDB.go
 * @Description: 这个文件是用来初始化数据库中的超级管理员和组织结构数据
 */
package main

import "go-sys/model"

var superAdmin = model.SysUser{
	UserName: "admin",
	Password: "admin",
	NickName: "超级管理员",
	Sex:      "男",
	Salt:     "salt",
	Email:    "admin@company.com",
	Phone:    "12345678901",
	// OrganizationId: 1,
	Status: 1,
}
var superAdminRole = model.SysRole{
	RoleName:  "超级管理员",
	Descption: "系统超级管理员",
	Status:    1,
}

var rootOrg = model.SysOrganization{
	Name:      "总公司",
	Sort:      1,
	Status:    1,
	Descption: "总公司",
}
var permission1 = model.SysPermission{
	PermissionId: 1,
	Resource:     "example-resource",
	Action:       "example-action-1",
}
var permission2 = model.SysPermission{
	PermissionId: 2,
	Resource:     "example-resource",
	Action:       "example-action-2",
}

// main is used to init super admin and role
func main() {
	model.InitDB()
	initrootOrg()
	initpermissions()
	initRole()
	initSuperAdmin()
	setSuperAdminRoleAndPermissions()
}

// initSuperAdmin is used to init super admin
func initSuperAdmin() {
	if err := model.GlobalDB.Create(&superAdmin).Error; err != nil {
		panic(err)
	}
}

// initRole is used to init role
func initRole() {
	if err := model.GlobalDB.Create(&superAdminRole).Error; err != nil {
		panic(err)
	}
}

// initrootOrg is used to set root org
func initrootOrg() {
	if err := model.GlobalDB.Create(&rootOrg).Error; err != nil {
		panic(err)
	}
}

// initpermissions is used to set permissions
func initpermissions() {

	// 创建权限
	if err := model.GlobalDB.Create(&permission1).Error; err != nil {
		panic(err)
	}
	if err := model.GlobalDB.Create(&permission2).Error; err != nil {
		panic(err)
	}
}

// setsuperAdminRoleAndPermissions is used to set super admin role and permissions
func setSuperAdminRoleAndPermissions() {
	// 查询所有权限
	var permissions []model.SysPermission
	if err := model.GlobalDB.Find(&permissions).Error; err != nil {
		panic(err)
	}

	// 查询超级管理员角色
	var superAdminRole model.SysRole
	if err := model.GlobalDB.Where("role_name = ?", "超级管理员").First(&superAdminRole).Error; err != nil {
		panic(err)
	}

	// 给超级管理员角色赋予权限
	if err := model.GlobalDB.Model(&superAdminRole).Association("Permissions").Append(&permissions); err != nil {
		panic(err)
	}

	// 查询超级管理员用户
	var superAdmin model.SysUser
	if err := model.GlobalDB.Where("user_name = ?", "admin").First(&superAdmin).Error; err != nil {
		panic(err)
	}

	// 将超级管理员角色赋予给超级管理员用户
	if err := model.GlobalDB.Model(&superAdmin).Association("Roles").Append(&superAdminRole); err != nil {
		panic(err)
	}

	// 查询组织结构
	var rootOrg model.SysOrganization
	if err := model.GlobalDB.Where("name = ?", "总公司").First(&rootOrg).Error; err != nil {
		panic(err)
	}

	// 将超级管理员用户赋予给组织结构
	if err := model.GlobalDB.Model(&rootOrg).Association("Users").Append(&superAdmin); err != nil {
		panic(err)
	}

}
