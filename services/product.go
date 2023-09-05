package services

import (
	"github.com/apirestgo/models"
	"github.com/apirestgo/repository"
)

func InsertProduct(product models.Product) (error, int) {
	id := repository.SetProducts(product)
	return nil, id
}

func UpsetProduct(product_id int, user_id int, quantity int) (error, models.Cart) {
	cart, err := repository.GetCartByUserId(user_id)
	product := repository.GetProductById(product_id)
	product_cart := models.Product_cart{
		Product:  product,
		Quantity: quantity,
	}
	if err != nil {
		if err.Error() == "cart not found" {
			cart = models.Cart{
				User_id: user_id,
				Product_cart: []models.Product_cart{
					product_cart,
				},
			}
			repository.SetCarts(cart)
			return nil, cart
		} else {
			return err, models.Cart{}
		}
	}
	for i, product_cart := range cart.Product_cart {
		if product_cart.Product.Id == product.Id {
			cart.Product_cart[i].Quantity += quantity
			repository.SetCartByUserId(user_id, cart)
			return nil, cart
		}
	}
	cart.Product_cart = append(cart.Product_cart, product_cart)
	repository.SetCartByUserId(user_id, cart)
	return nil, cart
}

func GetCart(user_id int) (models.Cart, error) {
	cart, err := repository.GetCartByUserId(user_id)
	if err != nil {
		return cart, err
	}
	return cart, nil
}

func GetOrder(user_id int) (models.Order, error) {
	cart, err := repository.GetCartByUserId(user_id)
	if err != nil {
		return models.Order{}, err
	}
	order := models.Order{
		Cart_id: cart.Id,
		Total: models.Total{
			Product:  0,
			Discount: 0,
			Shipping: 0,
			Order:    0,
		},
	}
	for _, product_cart := range cart.Product_cart {
		order.Total.Product += product_cart.Product.Price * float32(product_cart.Quantity)
	}
	if order.Total.Product > 100 {
		order.Total.Discount = order.Total.Product * 0.1
	}
	if order.Total.Product > 50 {
		order.Total.Shipping = 5
	}
	order.Total.Order = order.Total.Product - order.Total.Discount + order.Total.Shipping
	repository.SetOrders(order)
	repository.DeleteCartByUserId(user_id)
	return order, nil
}
