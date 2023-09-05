package models

type Cart struct {
	Id           int
	User_id      int
	Product_cart []Product_cart
}
type Product_cart struct {
	Id       int
	Product  Product
	Quantity int
}

type Product struct {
	Id       int
	Name     string
	Category string
	Price    float32
}

type Order struct {
	Id      int
	Cart_id int
	Total   Total
}

type Total struct {
	Product  float32
	Discount float32
	Shipping float32
	Order    float32
}

type Product_cart_dto struct {
	Quantity int
	Id       int
}

type Product_dto struct {
	Name     string
	Category string
	Price    float32
}

type User struct {
	Id       int
	Username string
	Password string
}
