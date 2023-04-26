package model

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

// Gorm关联关系配置
func configDBRelations() error {
	return GlobalDB.AutoMigrate(
		&SysUser{},
		&SysRole{},
		&SysPermission{},
		&SysOrganization{},
	)
}

func InitDB() error {
	viper.SetConfigFile("./config/dbConfig.yaml")
	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	// 从配置文件中获取数据库配置信息
	dbConfig := viper.GetStringMapString("database")

	// 格式化dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		dbConfig["user"],
		dbConfig["password"],
		dbConfig["host"],
		dbConfig["port"],
		dbConfig["dbname"],
		dbConfig["charset"],
		dbConfig["loc"],
	)

	// 上面的 dsn 中的 user、password、database 分别替换为自己的数据库用户名、密码、数据库名
	GlobalDB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                           dsn,
		DefaultStringSize:             256,
		DisableDatetimePrecision:      true,
		DontSupportRenameIndex:        true,
		DontSupportRenameColumn:       true,
		DontSupportNullAsDefaultValue: false,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return configDBRelations()
}
