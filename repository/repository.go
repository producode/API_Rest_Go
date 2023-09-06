package repository

import (
	"github.com/apirestgo/models"

	"errors"
)

var products []models.Product
var carts []models.Cart
var orders []models.Order
var users []models.User

func GetCartByUserId(user_id int) (models.Cart, error) {
	for _, cart := range carts {
		if cart.User_id == user_id {
			return cart, nil
		}
	}
	return models.Cart{}, errors.New("cart not found")
}

func InsertCart(cart models.Cart) (models.Cart, error) {
	cart.Id = len(carts) + 1
	carts = append(carts, cart)
	return cart, nil
}

func ModifyCart(cart models.Cart) (models.Cart, error) {
	for i, cart := range carts {
		if cart.User_id == cart.User_id {
			carts[i] = cart
			return cart, nil
		}
	}
	return cart, errors.New("cart not found")
}

func UpsertProductCart(product_cart models.Product_cart, user_id int) (models.Cart, error) {
	cart, err := GetCartByUserId(user_id)
	if err != nil {
		return cart, err
	}
	for i, product_cart := range cart.Product_cart {
		if product_cart.Product.Id == product_cart.Product.Id {
			cart.Product_cart[i] = product_cart
			return cart, nil
		}
	}
	cart.Product_cart = append(cart.Product_cart, product_cart)
	return cart, nil
}

func GetOrderByCartId(cart_id int) (models.Order, error) {
	for _, order := range orders {
		if order.Cart_id == cart_id {
			return order, nil
		}
	}
	return models.Order{}, errors.New("order not found")
}

func InsertOrder(order models.Order) (models.Order, error) {
	orders = append(orders, order)
	return order, nil
}

func InsertProduct(product models.Product) (models.Product, error) {
	product.Id = len(products) + 1
	products = append(products, product)
	return product, nil
}

func GetProductById(product_id int) (models.Product, error) {
	for _, product := range products {
		if product.Id == product_id {
			return product, nil
		}
	}
	return models.Product{}, errors.New("product not found")
}
