/*
 * @Author: Profigogogogo wf851128@gmail.com
 * @Date: 2023-04-26 21:13:43
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2023-04-27 11:40:30
 * @FilePath: /go-sys/model/sys_user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import (
	"go-sys/common/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents the user table and json data
type SysUser struct {
	UserId   uint       `gorm:"primaryKey;autoIncrement;comment:用户ID" json:"userId"` //用户ID
	UserName string     `gorm:"size:64;comment:用户名" json:"username"`                 //用户名
	Password string     `gorm:"size:128;comment:密码" json:"-"`                        //密码
	NickName string     `gorm:"size:128;comment:昵称" json:"nickName"`                 //昵称
	Sex      string     `gorm:"size:255;comment:性别" json:"sex"`                      //性别
	Salt     string     `gorm:"size:255;comment:加盐" json:"-"`                        //加盐
	Email    string     `gorm:"size:128;comment:邮箱" json:"email"`                    //邮箱
	Phone    string     `gorm:"size:11;comment:手机号" json:"phone"`                    //手机号
	Status   int        `gorm:"size:4;comment:状态" json:"status"`                     //状态
	Roles    []*SysRole `gorm:"many2many:user_role" json:"sys_role" `                //	角色关联
	// OrganizationId uint             `gorm:"comment:组织代码" json:"orgId"`                           //编码
	// Organization *SysOrganization `gorm:"foreignKey:OrganizationId"` //所属组织
	model.ModelTime
	model.ControlBy
}

// setting table name: sys_user
func (SysUser) TableName() string {
	return "sys_user"
}

// get user info
func (sys *SysUser) Generate() model.ActiveRecord {
	info := *sys
	return &info
}

// get user id
func (sys *SysUser) GetId() interface{} {
	return sys.UserId
}

// encrypt password
func (sys *SysUser) Encrypt() (err error) {
	if sys.Password == "" {
		return
	}
	//  生成加盐密码
	var hash []byte
	saltBytes := []byte(sys.Salt)
	passwordBytes := []byte(sys.Password)
	combinedBytes := append(passwordBytes, saltBytes...)
	// 将密码进行加密
	if hash, err = bcrypt.GenerateFromPassword([]byte(combinedBytes), bcrypt.DefaultCost); err != nil {
		return
	} else {
		sys.Password = string(hash)
		return
	}
}

// BeforeCreate 方法，当用户创建的时候，对密码进行加密。
func (sys *SysUser) BeforeCreate(_ *gorm.DB) error {
	return sys.Encrypt()
}

// BeforeUpdate 方法，如果用户修改了密码，那么就对密码进行加密，如果没有修改密码，那么就不对密码进行加密。
func (sys *SysUser) BeforeUpdate(_ *gorm.DB) error {
	var err error
	if sys.Password != "" {
		err = sys.Encrypt()
	}
	return err
}

// 1. 在查询数据库时，gorm 会自动调用 AfterFind 函数，将查询到的数据赋值给 SysUser 对象

// func (e *SysUser) AfterFind(_ *gorm.DB) error {
// 	e.Roles = []*SysRole{e.Roles}
// 	return nil
// }
