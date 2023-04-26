/*
 * @Author: Profigogogogo wf851128@gmail.com
 * @Date: 2023-04-26 21:13:21
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2023-04-26 23:29:27
 * @FilePath: /go-sys/model/sys_permission.go
 * @Description:权限配置
 */

package model

import "go-sys/common/model"

// Permission represents the permission table and json data
type SysPermission struct {
	PermissionId uint       `gorm:"primaryKey;comment:权限id" json:"permissionId"` //权限id
	Resource     string     `gorm:"size:128;comment:资源" json:"resource"`         //资源
	Action       string     `gorm:"size:128;comment:操作" json:"action"`           //操作
	Roles        []*SysRole `gorm:"many2many:role_permissions;"`                 //角色权限关联
	model.ControlBy
	model.ModelTime
}

// setting table name: sys_permission
func (SysPermission) TableName() string {
	return "sys_permission"
}

// get permission info
func (sys *SysPermission) Generate() model.ActiveRecord {
	info := *sys
	return &info
}

// get permission id
func (sys *SysPermission) GetId() interface{} {
	return sys.PermissionId
}
