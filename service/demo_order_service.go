package service

import (
	"gobigfirst/model"
	"strconv"
)

func Insert(demoOrder *model.DemoOrder) error {
	return model.GetDB().Create(demoOrder).Error
}

func Update(demoOrder *model.DemoOrder) error {
	return model.GetDB().Model(demoOrder).Updates(map[string]interface{}{
		"amount":   demoOrder.Amount,
		"status":   demoOrder.Status,
		"file_url": demoOrder.FileUrl,
	}).Error
}

func GetDemoOrderById(id string)(*model.DemoOrder, error) {
	pid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	var demoOrder model.DemoOrder
	err = model.GetDB().First(&demoOrder, "id = ?", pid).Error
	return &demoOrder, err
}


func GetDemoOrderList(userName string)([]*model.DemoOrder, error){
	var demoOrders []*model.DemoOrder
	//err := DB.Offset(from).Limit(size).Find(&demoOrders).Order("amount DESC").Order("created_at DESC").Error
	err := model.GetDB().Where("user_name like ?", "%"+userName+"%").Find(&demoOrders).Order("amount DESC").Order("created_at DESC").Error
	if err != nil {
		return nil, err
	}
	return demoOrders, err
}
