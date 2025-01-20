package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"orderservice/models"
	"orderservice/rabbitmq"
	"orderservice/repository"
)

func CreateOrderById(ctx *gin.Context) {

	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	id := order.UserID
	geturi := fmt.Sprintf("http://localhost:8080/user/user/user/%v", id)
	resp, err := http.Get(geturi)
	if err != nil {
		fmt.Println("cant get response")
		return
	}
	if resp.StatusCode != http.StatusOK {
		geturi1 := fmt.Sprintf("http://localhost:8080/user/register")
		resp2, err2 := http.Post(geturi1, "application/json", resp.Body)
		if err2 != nil {
			fmt.Println("cant create customer")
		}
		if resp2.StatusCode != http.StatusCreated {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not create customer"})
			return

		}

		//ctx.JSON(http.StatusUnauthorized, gin.H{"error": "customer does not exist"})

	}
	id2 := order.ProductID
	geturi2 := fmt.Sprintf("http://localhost:8080/user/products/get-product/%v", id2)
	resp, err = http.Get(geturi2)
	if err != nil {
		fmt.Println("cant get response from second uri")
		return
	}
	if resp.StatusCode != http.StatusOK {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "product does not exist"})
		return

	}

	_, err2 := repository.CreateOrder(order)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"order": "created"})

}
func GetAllOrders(ctx *gin.Context) {
	//var orders []models.order
	orders, err := repository.FindAllOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, orders)
}

func GetOrderByid(ctx *gin.Context) {
	id := ctx.Param("id")
	order, _ := repository.GetOrder(id)
	if order == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
	}
	ctx.JSON(http.StatusOK, gin.H{"order": order})

}
func BulkOrders(ctx *gin.Context) {

	var orders []models.Order
	if err := ctx.ShouldBindJSON(&orders); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	//var order models.Order
	for _, o := range orders {
		fmt.Printf("%+v\n", o)
		orderdata, err := json.Marshal(o)
		if err != nil {
			//ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			fmt.Println("cannot marshal order")
			return
		}
		fmt.Println("yoyo")
		err1 := rabbitmq.Channel.Publish(
			"order_exchange",
			"order_queue",
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        orderdata,
			},
		)
		fmt.Println("yeyey")
		if err1 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enqueue order"})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"order": o})

	}

}

func UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	var updatedOrder bson.M
	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := repository.UpdateByID(id, bson.M{"$set": updatedOrder}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

func DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := repository.DeleteByID(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "order deleted"})

}
