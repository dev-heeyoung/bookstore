package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:mysqlheeyoung!@/Books?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	return db
}
