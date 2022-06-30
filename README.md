# Deployments

## Connect to API Infinity

```bash
https://infinitysport.site
```
This is link development access by swagger
[Link](https://app.swaggerhub.com/apis-docs/Alfin7007/infinitysport/1.0.1#/)

## Endpoint to Access API

### USER

Access for get data, registration, login users

```bash
https://infinitysport.site/users

#example:
GET  https://infinitysport.site/users
POST https://infinitysport.site/users
```

**POST** method can be used for login and registration access

Access for update and delete

```bash
https://infinitysport.site/users/id

#example:
PUT https://infinitysport.site/users/id
DELETE https://infinitysport.site/users/id
```

### PRODUCT

Access for get data product, and add product

```bash
https://infinitysport.site/products

#example:
GET  https://infinitysport.site/products
POST https://infinitysport.site/products
```

Access for update and delete

```bash
https://infinitysport.site/products/id

#example:
PUT https://infinitysport.site/products/id
DELETE https://infinitysport.site/products/id
```

### ORDER

Access for get data order, and checkout order

```bash
https://infinitysport.site/orders

#example:
GET  https://infinitysport.site/orders
POST https://infinitysport.site/orders
```
for this is order, you can be used POST method for access to checkout 

Access for confirmation

```bash
https://infinitysport.site/orders/id

#example:
POST https://infinitysport.site/orders/id/confirm
```
this is method used for client to confirmation to payments success

Access for cancel orders

```bash
https://infinitysport.site/orders/id

#example:
POST https://infinitysport.site/orders/id/canceled
```

this is method used for client to cancelled orders

### CART

this is feature for waiting list order customer

```bash
https://infinitysport.site/carts

#example:
GET  https://infinitysport.site/carts
POST https://infinitysport.site/carts
```

update cart customer, you can be used method

```bash
https://infinitysport.site/carts/id

PUT    https://infinitysport.site/carts/id
DELETE https://infinitysport.site/carts/id
```

you can manipulation list carts, by PUT and DELETE method

## Feature List InfinitySport API
### USER
* Register
* Login
* Authorized 
* List Account
* Delete Account
* Update Account

### ORDER
* Order List
* Checkout payment
* Cancel Order

### CART
* Cart List
* Manipulation Cart List

### PRODUCT
* Product List management
* Add Product
* Update Product
* Delete Product
