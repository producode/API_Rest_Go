package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Cart struct {
	Id           int
	User_id      int
	Product_cart []Product_cart
}
type Product_cart struct {
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
	Cart_id int
	Total   Total
}

type Total struct {
	Product  float32
	Discount float32
	Shipping float32
	Order    float32
}

type User struct {
	Id       int
	Username string
	Password string
}

type UserClaims struct {
	User_id int `json:"user_id"`
	jwt.StandardClaims
}

const (
	Coffee     = "Coffee"
	Equipment  = "Equipment"
	Accesories = "Accesories"
)
