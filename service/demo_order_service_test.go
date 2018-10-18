package service

import (
	"fmt"
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

func Test_GetDemoOrderById(t *testing.T) {
	id := "1"
	demoOrder,err := GetDemoOrderById(id)
	fmt.Println(demoOrder, err)
}

func Test_GetDemoOrderList(t *testing.T)  {
	userName := "lis"
	demoOrders, err := GetDemoOrderList(userName)
	if err == nil {

		for i := 0; i< len(demoOrders); i++ {
			fmt.Println(demoOrders[i])
		}
	}
}

func Test_string(t *testing.T)  {
	for i := 0; i < 10; i++ {
		fmt.Println(string(i))
	}
}