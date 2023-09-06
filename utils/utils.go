package utils

import (
	"github.com/apirestgo/models"
	"github.com/dgrijalva/jwt-go"
)

func Calculate_discount(products []models.Product_cart, total float32) float32 {
	var discount float32 = 0
	var total_discount_apply bool = false
	for _, product_cart := range products {
		if product_cart.Product.Category == models.Coffee && product_cart.Quantity >= 2 {
			discount += (float32(product_cart.Quantity) * product_cart.Product.Price)
		}
		if product_cart.Product.Category == models.Accesories && (float32(product_cart.Quantity)*product_cart.Product.Price) > 70 {
			total_discount_apply = true
		}
	}
	if total_discount_apply {
		discount += (total * 0.1)
	}
	return discount
}

func Calculate_shipping(products []models.Product_cart, shipping float32) float32 {
	for _, product := range products {
		if product.Product.Category == models.Equipment && product.Quantity >= 3 {
			shipping = 0
		}
	}
	return shipping
}

func Generate_JWT(user models.User) string {
	claims := models.UserClaims{
		User_id: user.Id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte("secret"))
	return signedToken
}
