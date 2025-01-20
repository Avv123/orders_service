package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"orderservice/config"
	"orderservice/models"
	"time"
)

var (
	OrderCollection *mongo.Collection
)

func InitOrderCollection() {
	OrderCollection = config.GetCollection("orders")
}
func CreateOrder(order models.Order) (*models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := OrderCollection.InsertOne(ctx, order)
	if err != nil {
		return nil, err

	}
	return &order, nil

}

func GetOrder(id string) (*models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var order models.Order
	err := OrderCollection.FindOne(ctx, bson.M{"id": id}).Decode(&order)
	if err != nil {
		return nil, err

	}
	return &order, nil

}
func FindAllOrders() ([]models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var orders []models.Order
	cursor, err := OrderCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var order models.Order
		err := cursor.Decode(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil

}
func UpdateByID(id string, updatedOrder bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := OrderCollection.UpdateOne(ctx, bson.M{"_id": id}, updatedOrder)
	return err
}

func DeleteByID(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := OrderCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
