package main

import (
	"github.com/gin-gonic/gin"

	"myGo/service"
)

func main() {
	r := gin.Default()
	pr := r.Group("/product")
	// 增删查改
	{
		pr.POST("/add", service.AddProduct)
		pr.PUT("/update", service.UpdateProduct)
		pr.GET("/id/:id", service.GetOne)
		pr.GET("", service.GetAll)
		pr.DELETE("/delete", service.DeleteProduct)
	}
	r.Run(":80")
}
