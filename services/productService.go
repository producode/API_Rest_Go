package services

import (
	"github.com/apirestgo/models"
	"github.com/apirestgo/repository"
	"github.com/apirestgo/utils"

	"errors"
	"regexp"
)

func InsertProduct(product models.Product) (error, models.Product) {
	var validCategory = regexp.MustCompile("^(?i)(Coffee|Equipment|Accesories)$")
	if product.Name == "" {
		return errors.New("name is required"), models.Product{}
	}
	if !validCategory.MatchString(product.Category) {
		return errors.New("not valid category "), models.Product{}
	}
	if product.Price <= 0 {
		return errors.New("price is required"), models.Product{}
	}
	product, err := repository.InsertProduct(product)
	if err != nil {
		return err, models.Product{}
	}
	return nil, product
}

func AddToCart(product_id int, user_id int) (error, models.Cart) {
	product, err := repository.GetProductById(product_id)
	if err != nil {
		return err, models.Cart{}
	}
	product_cart := models.Product_cart{
		Product:  product,
		Quantity: 1,
	}
	cart, err := repository.GetCartByUserId(user_id)
	if err != nil {
		if err.Error() == "cart not found" {
			cart = models.Cart{
				User_id: user_id,
				Product_cart: []models.Product_cart{
					product_cart,
				},
			}
			cart, err := repository.InsertCart(cart)
			if err != nil {
				return err, models.Cart{}
			}
			return nil, cart
		} else {
			return err, models.Cart{}
		}
	}
	repository.GetCartByUserId(user_id)

	cart.Product_cart = append(cart.Product_cart, product_cart)
	cart, err = repository.UpsertProductCart(product_cart, user_id)
	if err != nil {
		return err, models.Cart{}
	}
	return nil, cart
}

func ModifyCartQuantity(product_id int, user_id int, quantity int) (error, models.Cart) {
	cart, err := repository.GetCartByUserId(user_id)
	if err != nil {
		return err, models.Cart{}
	}
	for i, product_cart := range cart.Product_cart {
		if product_cart.Product.Id == product_id {
			cart.Product_cart[i].Quantity = quantity
			cart, err = repository.ModifyCart(cart)
			if err != nil {
				return err, models.Cart{}
			}
			return nil, cart
		}
	}
	return errors.New("product not found"), models.Cart{}
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
			Shipping: 5,
			Order:    0,
		},
	}
	for _, product_cart := range cart.Product_cart {
		order.Total.Product += product_cart.Product.Price * float32(product_cart.Quantity)
	}
	order.Total.Discount = utils.Calculate_discount(cart.Product_cart, order.Total.Product)
	order.Total.Shipping = utils.Calculate_shipping(cart.Product_cart, order.Total.Shipping)
	order.Total.Order = order.Total.Product - order.Total.Discount + order.Total.Shipping
	order, err = repository.InsertOrder(order)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}
