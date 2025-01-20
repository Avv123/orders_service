package routes

import (
	"github.com/gin-gonic/gin"
	"orderservice/handlers"
)

func SetupRoutes(router *gin.Engine) {

	//router := gin.Default()
	api := router.Group("/orders")
	api.POST("/create-order", handlers.CreateOrderById)
	api.POST("/bulk-orders", handlers.BulkOrders)
	api.GET("/get-product/:id", handlers.GetOrderByid)
	api.GET("/getallproducts", handlers.GetAllOrders)
	api.PUT("/updateproduct", handlers.UpdateOrder)
	api.DELETE("/delete-product/:id", handlers.DeleteOrder)

}
