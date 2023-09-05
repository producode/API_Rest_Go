package utils

import "github.com/apirestgo/models"

func Calculate_discount(products []models.Product_cart, total float32) float32 {
	var discount float32 = 0
	var total_discount_apply bool = false
	for _, product := range products {
		if product.Product.Category == models.Coffee {
			quantity_rule(product, &discount, 2, 1)
		}
		if product.Product.Category == models.Accesories && (float32(product.Quantity)*product.Product.Price) > 70 {
			total_discount_apply = true
		}
	}
	if total_discount_apply {
		percentage_rule(&discount, total, 0.1)
	}
	return discount
}

func Calculate_shipping(products []models.Product_cart, shipping float32) float32 {
	for _, product := range products {
		if product.Product.Category == models.Equipment && product.Quantity >= 3 {
			free_shipping_rule(&shipping)
		}
	}
	return shipping
}

func quantity_rule(product_cart models.Product_cart, discount *float32, quantity_limit int, quantity_discount int) {
	if product_cart.Product.Category == models.Coffee && product_cart.Quantity >= quantity_limit {
		*discount += (float32(quantity_discount) * product_cart.Product.Price)
	}
}

func free_shipping_rule(shipping *float32) {
	*shipping = 0
}

func percentage_rule(discount *float32, total float32, percentage float32) {
	*discount += (total * percentage)
}
