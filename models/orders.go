package models

type Order struct {
	ID        string `json:"id" bson:"_id"`
	Suk       string `json:"suk" bson:"suk"`
	Name      string `json:"name" bson:"name"`
	UserID    string `json:"userid" bson:"userid"`
	ProductID string `json:"productid" bson:"productid"`
	Quantity  int    `json:"quantity" bson:"quantity"`
}
