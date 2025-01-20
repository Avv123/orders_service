package consumers

import (
	"encoding/json"
	"fmt"
	//"github.com/gin-gonic/gin"
	"log"
	//"net/http"
	"orderservice/models"
	"orderservice/repository"

	// "net/smtp"
	"orderservice/rabbitmq"
)

type Order struct {
	ID        string `json:"id" bson:"_id"`
	Suk       string `json:"suk" bson:"suk"`
	Name      string `json:"name" bson:"name"`
	UserID    string `json:"userid" bson:"userid"`
	ProductID string `json:"productid" bson:"productid"`
	Quantity  int    `json:"quantity" bson:"quantity"`
}

func StartEmailConsumer() {
	rabbitmq.InitRabbitMQ()

	//defer rabbitmq.Conn.Close()
	//defer rabbitmq.Channel.Close()
	fmt.Println("Connected to RabbitMQ")

	q, err := rabbitmq.Channel.QueueDeclare(
		"order_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}
	fmt.Printf("yaha tak aaya kya")

	err = rabbitmq.Channel.QueueBind(
		q.Name,
		"order_queue",
		"order_exchange",
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind queue: %s", err)
	}

	orders, err := rabbitmq.Channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	go func() {
		for d := range orders {
			var o Order
			err := json.Unmarshal(d.Body, &o)

			if err != nil {
				log.Printf("cant unmaeshall")
				continue
			}
			result, err := repository.CreateOrder(models.Order(o))
			if err != nil {
				log.Printf("cant create order")
			}
			//ctx.JSON(http.StatusOK, gin.H{"order": result})

			fmt.Println(result)
		}
	}()
	log.Printf("Waiting for messages. To exit press CTRL+C")
	select {}

}
