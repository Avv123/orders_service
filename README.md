### Setup for the project
```
go get github.com/gin-gonic/gin
go get github.com/joho/godotenv
go.mongodb.org/mongo-driver/mongo
go get github.com/streadway/amqp
go run main.go


```

### Endpoint: orders/create-order

### Input:

 curl --location 'http://localhost:8081/orders/create-order' \ers/create-order' \
--header 'Content-Type: application/json' \
--data '   {
    "id":"order11",
    "suk":"26gb",
    "name":"iphone15",
    "userid":"5",
    "productid":"e",
    "quantity":45
    }
'


### Output:


{"order":"created"}
![image](https://github.com/user-attachments/assets/1d25d131-86fb-447b-8f09-3090fd358981)

### Endpoint: orders/bulkorder
### Input: 
curl --location 'http://localhost:8081/orders/bulk-orders' --header 'Content-Type: application/json' --data '[
    
        {
    "id":"order4",
    "suk":"256gb",
    "name":"iphone13",
    "userid":"2",
    "productid":"e",
    "quantity":45
        },
        {
        "id":"order3",
    "suk":"26gb",
'   ]   }tity":45e",,

### Output:


{"order":{"id":"order4","suk":"256gb","name":"iphone13","userid":"2","productid":"e","quantity":45}}{"order":{"id":"order3","suk":"26gb","name":"iphone3","userid":"2","productid":"e","quantity":45}}

![image](https://github.com/user-attachments/assets/9ff45c7a-cd92-4718-bdfc-8b5da2bb1563)

### Routes on console:

![image](https://github.com/user-attachments/assets/03ba8a19-9e74-4038-aecf-f659d679c416)




