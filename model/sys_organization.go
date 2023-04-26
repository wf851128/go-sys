/*
 * @Author: Profigogogogo wf851128@gmail.com
 * @Date: 2023-04-26 21:13:09
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2023-04-27 00:02:56
 * @FilePath: /go-sys/model/sys_organization.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import "go-sys/common/model"

// Organization represents the organization table and json data

type SysOrganization struct {
	OrganizationId uint              `gorm:"primaryKey;autoIncrement;comment:组织ID" json:"orgID"` //组织ID
	ParentID       uint              `gorm:"size:4;comment:父级ID" json:"parentID"`                //父级ID
	Name           string            `gorm:"size:128;comment:组织名称" json:"orgName"`               //组织名称
	Sort           int               `gorm:"size:4;comment:组织排序" json:"sort"`                    //组织排序
	Leader         string            `gorm:"size:128;comment:负责人" json:"leader"`                 //负责人
	Phone          string            `gorm:"size:11;comment:手机" json:"phone"`                    //手机
	Email          string            `gorm:"size:64;comment:邮箱" json:"email"`                    //邮箱
	Status         int               `gorm:"size:4;comment:状态" json:"status"`                    //状态
	Descption      string            `gorm:"size:128;comment:组织描述" json:"descption"`             // 组织描述
	Children       []SysOrganization `gorm:"-" json:"children"`                                  // 子级组织
	Users          []SysUser         `gorm:"many2many:organization_user;"`                       // 岗位用户关联
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
	return sys.OrganizationId
}
