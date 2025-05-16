#!/bin/bash

set -e

BASE_URL="http://localhost:8080/v1/inventory"
STATS_URL="http://localhost:8080/v1/statistics"
TEST_USER="user123"

echo ">>> Creating category..."
category_id=$(curl -s -X POST "$BASE_URL/category" \
  -H "Content-Type: application/json" \
  -d '{"name": "gadgets"}' | jq -r .id)
echo "Created category: $category_id"

echo ">>> Listing categories..."
curl -s "$BASE_URL/categories" | jq .

echo ">>> Creating product..."
product_id=$(curl -s -X POST "$BASE_URL/product" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "apple",
    "price": 25.99,
    "category": "'"$category_id"'",
    "stock": 1234
  }' | jq -r .id)
echo "Created product: $product_id"

echo ">>> Getting product by ID..."
curl -s "$BASE_URL/product/$product_id" | jq .

echo ">>> Updating product..."
curl -s -X PUT "$BASE_URL/product" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "'"$product_id"'",
    "name": "banana",
    "price": 24.25,
    "category": "'"$category_id"'",
    "stock": 324
  }' | jq .

echo ">>> Listing all products..."
curl -s "$BASE_URL/products" | jq .

echo ">>> Creating order..."
order_id=$(curl -s -X POST http://localhost:8080/v1/orders/ \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "'"$TEST_USER"'",
    "items": [
      { "product_id": "'"$product_id"'", "quantity": 1 }
    ]
  }' | jq -r .id)
echo "Created order: $order_id"

echo ">>> Checking user order statistics..."
curl -s "$STATS_URL/user/$TEST_USER/orders" | jq .

echo ">>> Checking global user statistics..."
curl -s "$STATS_URL/users" | jq .

echo ">>> Deleting product..."
curl -s -X DELETE "$BASE_URL/product/$product_id"
echo ""

echo ">>> Deleting category..."
curl -s -X DELETE "$BASE_URL/category/$category_id"
echo ""

echo ">>> Verifying product deletion..."
curl -s "$BASE_URL/product/$product_id" || echo "Product correctly deleted"
echo ""

echo ">>> Verifying category deletion..."
curl -s "$BASE_URL/category/$category_id" || echo "Category correctly deleted"
echo ""

echo ">>> Verifying user statistics still exists..."
curl -s "$STATS_URL/user/$TEST_USER/orders" | jq .

echo "All tests completed."
