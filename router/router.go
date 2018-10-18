package main

import (
	"github.com/gin-gonic/gin"
	"gobigfirst/handler"
)

func main()  {
	r := gin.Default()


	r.GET("/demoOrder", handler.GetDemoOrder)
	r.POST("/demoOrder", handler.CreateDemoOrder)
	r.PUT("/demoOrder", handler.UpdateDemoOrder)

	r.GET("/demoOrders")


	r.Run(":8080")
}
