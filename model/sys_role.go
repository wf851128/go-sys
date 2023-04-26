/*
 * @Author: Profigogogogo wf851128@gmail.com
 * @Date: 2023-04-26 21:13:29
 * @LastEditors: Profigogogogo wf851128@gmail.com
 * @LastEditTime: 2023-04-26 23:15:28
 * @FilePath: /go-sys/model/sys_role.go
 * @Description: 这个文件是用来保存角色的
 */
package model

import "go-sys/common/model"

// Role represents the role table
type SysRole struct {
	RoleId      int              `gorm:"primaryKey;autoIncrement;comment:角色ID" json:"roleId"` // 角色ID
	RoleName    string           `gorm:"size:128;comment:角色名称" json:"roleName"`               // 角色名称
	Descption   string           `gorm:"size:128;comment:描述" json:"descption"`                // 角色描述
	Status      int              `gorm:"size:4;comment:状态" json:"status"`                     // 状态
	Permissions []*SysPermission `gorm:"many2many:role_permissions;"`                         // 角色权限关联
	Users       []*SysUser       `gorm:"many2many:user_role"`                                 // 角色用户关联
	model.ControlBy
	model.ModelTime
}

// setting table name: sys_role
func (SysRole) TableName() string {
	return "sys_role"
}

// get Role info
func (sys *SysRole) Generate() model.ActiveRecord {
	info := *sys
	return &info
}

// get Role id
func (sys *SysRole) GetId() interface{} {
	return sys.RoleId
}
