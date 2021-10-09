# walk-api 
walk-API is consumed by walk-shoe stores that handle JSON format for request and response.

# Register and Sign In
Note: I have not implemented password hashing. The passwords are store in plain text.


To register, the customer needs to provide a name, email, and password.\
	POST /v1/register

To Sign  In, the customer needs an email and password.\
	POST /v1/signin

# Categories
To get product information by specific categories. 
 
Women.\
	Boots\
	GET /v1/categorie?cat=womens-boots\
	Heels\
	GET /v1/categorie?cat=heels

Men.\
	Boots\
	GET /v1/categorie?cat=mens-boots\
	Sport\
	GET /v1/categorie?cat=mens-sport


# Product 
To get more information of a product that was requested by product id.\
	GET /v1/product?product-id=1234	

# Promotions
The products on promotions can be customized as the store need them.
NOTE: All changes on products promotion are done on the database.\
	GET /v1/promo

All product that are in promotion or discount by product id.\
	GET /v1/promo?product-id=1234



# Orders
To make a purchase a customer needs a card, cv, and amount to proceed.\
	POST /v1/orders

To get a customer order or orders user id or customer id.\
	GET /v1/orders?=user-id=2

Refund and delete
Delete and refund a purchase id is required.\
	DELETE /v1/orders?del-refund=2 
