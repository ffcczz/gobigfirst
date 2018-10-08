package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	db2 "gobigfirst/db"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DemoOrder struct {
	gorm.Model
	OrderId  string
	UserName string
	Amount   float64
	Status   string
	FileUrl  string
}

var DB *gorm.DB



func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local" , db2.USERNAME, db2.PASSWORD, "test" ))
	//db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local" , db2.USERNAME, db2.PASSWORD, db2.HOSTNAME, db2.PORT , "test" ))
	//db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8&parseTime=True&loc=Local" )

	if err == nil {
		DB = db
		db.AutoMigrate(&DemoOrder{})
		return db, err

	}
	fmt.Println("获取数据库连接失败:", err)

	return nil, err
}
func GetDB() *gorm.DB {
	if DB == nil {
		InitDB()
	}
	return DB
}




