package model

import "go-sys/common/model"

// Role represents the role table
type SysRole struct {
	RoleId      int              `json:"roleId" gorm:"primaryKey;autoIncrementcomment:角色ID"` // 角色ID
	RoleName    string           `json:"roleName" gorm:"size:128;comment:角色名称"`              // 角色名称
	Descption   string           `json:"descption" gorm:"size:128;comment:描述"`               // 角色描述
	Status      int              `json:"status" gorm:"size:4;comment:状态"`                    // 状态
	Permissions []*SysPermission `gorm:"many2many:role_permissions;"`                        // 角色权限关联
	Users       []*SysUser       `gorm:"many2many:user_role;"`                               // 角色用户关联
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
