/*
 * @Author: Profigogogogo wf851128@gmail.com
 * @Date: 2023-04-26 16:33:24
 * @LastEditors: Profigogogogo wf851128@gmail.com
 * @LastEditTime: 2023-04-26 20:48:01
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
	OrgID:    1,
	Status:   1,
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

// main is used to init super admin and role
func main() {
	model.InitDB()
	// initSuperAdmin()
	// initRole()
	// initpermissions()
	// initrootOrg()
	// setSuperAdminRoleAndPermissions()
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
	permission1 := model.SysPermission{
		PermissionID: 1,
		Resource:     "example-resource",
		Action:       "example-action-1",
	}
	permission2 := model.SysPermission{
		PermissionID: 2,
		Resource:     "example-resource",
		Action:       "example-action-2",
	}

	// 创建权限
	if err := model.GlobalDB.Create(&permission1).Error; err != nil {
		panic(err)
	}
	if err := model.GlobalDB.Create(&permission2).Error; err != nil {
		panic(err)
	}
}

// setSuperAdminRole is used to set super admin role and permissions
func setSuperAdminRoleAndPermissions() {
	// 添加全部权限到超级管理员角色
	var permissions []model.SysPermission
	if err := model.GlobalDB.Find(&permissions).Error; err != nil {
		panic(err)
	}

	if err := model.GlobalDB.Model(&superAdminRole).
		Association("Permissions").
		Append(permissions); err != nil {
		panic(err)
	}

	// 将超级管理员关联到角色
	superAdminRole.Users = append(superAdminRole.Users, &superAdmin)
	if err := model.GlobalDB.Save(&superAdminRole).Error; err != nil {
		panic(err)
	}
}