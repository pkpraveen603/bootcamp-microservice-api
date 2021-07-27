# bootcamp-microservice-api
Day 4-5 assignment

This api runs on https://localhost:8080/

### The following requests can be made via this microservice (With the expected result):

#### Customer-level commands
1. GET "/customer-api/user" -> Get all customers in the table customer
2. POST "/customer-api/user" -> Create a new customer
3. PATCH "/customer-api/user/:id" -> Update an existing customer
4. GET "/customer-api/user/:id" -> Get the customer details with a particulat id
5. DELETE "/customer-api/user/:id" -> Delete a particular customer

#### Retalier-level commands
1. GET "/retailer-api/product/" Get all products
2. POST "/reatiler-api/product/" Create a new product
3. PATCH "/retailer-api/product/:id" Update an existing product
4. DELETE "/retailer-api/product/:id" Delete an existing product
5. GET "/retailer-api/product/:id" Get product by given id
6. PATCH "/retailer-api/order/:id" Update the order status to processed (done by retailer)

#### Commands for dealing with orders
1. GET "/order-api/order" Get all transations history
2. POST "/order-api/order" Create a new order
3. GET "/order-api/order/:id" Get the transaction history of customer with asked id in the query

