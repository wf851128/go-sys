package model

import "go-sys/common/model"

// Permission represents the permission table and json data
type SysPermission struct {
	PermissionID uint       `json:"permissionId" gorm:"primaryKey;comment:权限id"` //权限id
	Resource     string     `json:"resource"  gorm:"size:128;comment:资源"`        //资源
	Action       string     `json:"action" gorm:"size:128;comment:操作"`           //操作
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
	return sys.PermissionID
}
