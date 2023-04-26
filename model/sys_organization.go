package model

import "go-sys/common/model"

// Organization represents the organization table and json data
type SysOrganization struct {
	OrgID     uint               `json:"orgID" gorm:"primaryKey;autoIncrement;comment:组织ID"` //组织ID
	ParentID  uint               `json:"parentID" gorm:"size:4;comment:父级ID"`                //父级ID
	Name      string             `json:"orgName" gorm:"size:128;comment:组织名称"`               //组织名称
	Sort      int                `json:"sort" gorm:"size:4;comment:组织排序"`                    //组织排序
	Leader    string             `json:"leader" gorm:"size:128;comment:负责人"`                 //负责人
	Phone     string             `json:"phone" gorm:"size:11;comment:手机"`                    //手机
	Email     string             `json:"email" gorm:"size:64;comment:邮箱"`                    //邮箱
	Status    int                `json:"status" gorm:"size:4;comment:状态"`                    //状态
	Descption string             `json:"descption" gorm:"size:128;comment:组织描述"`             // 组织描述
	Children  []*SysOrganization `json:"children" gorm:"-"`                                  // 子级组织
	Users     []*SysUser         `gorm:"many2many:user_organization;"`                       // 岗位用户关联

	model.ControlBy
	model.ModelTime
}

// setting table name: sys_organization
func (SysOrganization) TableName() string {
	return "sys_organization"
}

// get organization info
func (sys *SysOrganization) Generate() model.ActiveRecord {
	info := *sys
	return &info
}

// get organization id
func (sys *SysOrganization) GetId() interface{} {
	return sys.OrgID
}
