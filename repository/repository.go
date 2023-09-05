package repository

import (
	"github.com/apirestgo/models"

	"errors"
)

var products []models.Product
var carts []models.Cart
var orders []models.Order
var users []models.User

func GetProducts() []models.Product {
	return products
}
func GetCarts() []models.Cart {
	return carts
}
func GetOrders() []models.Order {
	return orders
}
func GetUsers() []models.User {
	return users
}
func SetProducts(product models.Product) int {
	product.Id = len(products) + 1
	products = append(products, product)
	return product.Id
}
func SetCarts(cart models.Cart) int {
	cart.Id = len(carts) + 1
	carts = append(carts, cart)
	return cart.Id
}
func SetOrders(order models.Order) {
	orders = append(orders, order)
}
func SetUsers(user models.User) {
	users = append(users, user)
}
func GetProductById(id int) (models.Product, error) {
	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}
	return models.Product{}, errors.New("product not found")
}
func GetCartById(id int) models.Cart {
	for _, cart := range carts {
		if cart.Id == id {
			return cart
		}
	}
	return models.Cart{}
}
func GetOrderById(id int) models.Order {
	for _, order := range orders {
		if order.Id == id {
			return order
		}
	}
	return models.Order{}
}
func GetCartByUserId(id int) (models.Cart, error) {
	for _, cart := range carts {
		if cart.User_id == id {
			return cart, nil
		}
	}
	return models.Cart{}, errors.New("cart not found")
}

func SetCartByUserId(id int, cart models.Cart) {
	for i, cart_find := range carts {
		if cart_find.User_id == id {
			carts[i] = cart
		}
	}
}

func DeleteCartByUserId(id int) {
	for i, cart := range carts {
		if cart.User_id == id {
			carts = append(carts[:i], carts[i+1:]...)
		}
	}
}
