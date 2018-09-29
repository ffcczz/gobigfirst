package model

import (
	"github.com/jinzhu/gorm"
	"strconv"
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
	db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8&parseTime=True&loc=Local")

	if err == nil {
		DB = db
		db.AutoMigrate(&DemoOrder{})
		return db, err

	}

	return nil, err
}

func (demoOrder *DemoOrder) Insert() error {
	return DB.Create(demoOrder).Error
}

func (demoOrder *DemoOrder) update() error {
	return DB.Model(demoOrder).Updates(map[string]interface{}{
		"amount":   demoOrder.Amount,
		"status":   demoOrder.Status,
		"file_url": demoOrder.FileUrl,
	}).Error
}

func GetDemoOrderById(id string)(*DemoOrder, error) {
	pid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	var demoOrder DemoOrder
	err = DB.First(&demoOrder, "id = ?", pid).Error
	return &demoOrder, err
}


func GetDemoOrderList(userName string)([]*DemoOrder, error){
	var demoOrders []*DemoOrder
	//err := DB.Offset(from).Limit(size).Find(&demoOrders).Order("amount DESC").Order("created_at DESC").Error
	err := DB.Where("user_name like ?", "%"+userName+"%").Find(&demoOrders).Order("amount DESC").Order("created_at DESC").Error
	if err != nil {
		return nil, err
	}
	return demoOrders, err
}



