package router

import (
	"github.com/gin-gonic/gin"
	"gobigfirst/handler"
)

func main()  {
	r := gin.Default()
	r.Group("/demoOrder")
	{

		r.GET("/", handler.GetDemoOrder)
		r.POST("/", handler.CreateDemoOrder)
		r.PUT("/", handler.UpdateDemoOrder)
	}
	r.Group("/demoOrders")
	{
		r.GET("/", handler.GetDemoOrderList)
	}

	r.Run(":8080")
}
