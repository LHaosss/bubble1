package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

//InitMysql 连接数据库
func InitMysql() (err error) {
	dsn := "root:293977@(172.21.229.89:3306)/bubble1?cahrset=utf8mb4&parseTime=True&loc=Local"
	//使用gorm语句连接数据库
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = DB.DB().Ping()
	return
}
