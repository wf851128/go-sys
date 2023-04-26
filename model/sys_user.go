package model

import (
	"go-sys/common/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents the user table and json data
type SysUser struct {
	UserId   uint       `json:"userId" gorm:"primaryKey;autoIncrement;comment:用户ID"` //用户ID
	UserName string     `json:"username" gorm:"size:64;comment:用户名"`                 //用户名
	Password string     `json:"-" gorm:"size:128;comment:密码"`                        //密码
	NickName string     `json:"nickName" gorm:"size:128;comment:昵称"`                 //昵称
	Sex      string     `json:"sex" gorm:"size:255;comment:性别"`                      //性别
	Salt     string     `json:"-" gorm:"size:255;comment:加盐"`                        //加盐
	Email    string     `json:"email" gorm:"size:128;comment:邮箱"`                    //邮箱
	Phone    string     `json:"phone" gorm:"size:11;comment:手机号"`                    //手机号
	OrgID    uint       `json:"orgID" gorm:"primaryKey;comment:编码"`                  //编码
	Status   int        `json:"status" gorm:"size:4;comment:状态"`                     //状态
	Roles    []*SysRole `json:"sysr_roles" gorm:"many2many:user_role;"`              //	角色关联
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

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(sys.Password), bcrypt.DefaultCost); err != nil {
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
// 2. 在调用 AfterFind 函数时，SysUser 对象的 DeptIds、PostIds、RoleIds 字段还没有赋值，所以需要在 AfterFind 函数中手动赋值
// 3. 在调用 AfterFind 函数时，SysUser 对象的 Dept、Post、Role 字段还没有赋值，所以需要在 AfterFind 函数中手动赋值
// func (e *SysUser) AfterFind(_ *gorm.DB) error {
// 	e.Roles = []int{e.Roles}
// 	return nil
// }
