package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gobigfirst/model"
)
import "gobigfirst/service"
func GetDemoOrder(c *gin.Context)  {
	orderId := c.Param("id")
	if orderId == "" {
		c.JSON(200, "参数不正确")
		return
	}
	demoOrder, err := service.GetDemoOrderById(orderId)
	if err != nil {
		c.JSON(200, err)
		return
	}

	c.JSON(200, demoOrder)
}

func CreateDemoOrder(c *gin.Context)  {
	var demoOrder *model.DemoOrder
	if err := c.BindJSON(demoOrder); err != nil {
		c.JSON(200, "参数错误")
		return
	}

	err := service.Insert(demoOrder)
	if err != nil {
		c.JSON(200, err)
		return
	}
	c.JSON(200, "插入成功")

}

func UpdateDemoOrder(c *gin.Context)  {
	var demoOrder *model.DemoOrder
	if err := c.BindJSON(demoOrder); err != nil {
		c.JSON(200, "参数错误")
		return
	}
	err := service.Update(demoOrder)
	if err != nil {
		c.JSON(200, err)
		return
	}
	c.JSON(200, "更新成功 ")
}

func GetDemoOrderList(c *gin.Context)  {
	userName := c.Param("userName")
	fmt.Println("需要查询的userName= ", userName)
	demoOrders, err := service.GetDemoOrderList(userName)
	if err != nil {
		c.JSON(200, err)
		return
	}
	c.JSON(200, demoOrders)
}
