package service

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"todo-app/config"
)

var Db *gorm.DB

//加载数据库
func LoadDb(){
	Db = ConnectDB()
	Db.LogMode(true)
}

//连接数据库
func ConnectDB() *gorm.DB{
	Db, err := gorm.Open("mysql", config.Conf.Db.Mysql)
	if nil != err {
		panic("failed to connect database")
	}
	return Db
}

//关闭数据库
func DisconnectDB(){
	if err := Db.Close(); nil != err{
		//关闭失败
	}
}