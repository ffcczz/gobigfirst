package service

import (
	"gobigfirst/model"
	"testing"
)

func Test_Insert(t *testing.T)  {
	/*
	OrderId  string
	UserName string
	Amount   float64
	Status   string
	FileUrl  string
	*/
	demoOrder := model.DemoOrder{
		OrderId: "first",
		UserName: "lisi",
		Amount: 123,
		Status: "test",
		FileUrl: "/usr/local/demoOrder/insert.png",
	}

	for i := 0; i < 10; i++ {
		demoOrder.ID = uint(i) + 2;
		demoOrder.OrderId = "first_" + string(i);
		Insert(&demoOrder)
	}
}

func Test_Update(t *testing.T)  {
	demoOrder := model.DemoOrder{
		Amount: 155,
		FileUrl: "/usr/local/demoOrder/update.png",
		Status: "update",

	}
	Update(&demoOrder)
}