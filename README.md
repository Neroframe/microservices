### Simple ecommerce microservices

To start services: 
docker compose up --build -d 

I used protoc cmd below:
protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. proto/file_name.proto




curl -X POST http://localhost:8080/v1/inventory/product \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Espresso",
    "price": 4.99,
    "category": "coffee",
    "stock": 50
  }'
  
curl -X POST http://localhost:8080/v1/inventory/category \
  -H "Content-Type: application/json" \
  -d '{ "name": "coffee" }'

curl -X POST http://localhost:8080/v1/orders/ \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user123",
    "items": [
      { "product_id": "PRODUCT_ID", "quantity": 2 }
    ]
  }'
  
curl -X POST http://localhost:8080/v1/payments/ \
  -H "Content-Type: application/json" \
  -d '{
    "order_id": "ORDER_ID",
    "amount": 9.98,
    "payment_method": "Credit Card"
  }'
