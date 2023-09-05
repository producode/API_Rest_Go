package main

import (
	"testing"

	"github.com/apirestgo/models"
	"github.com/apirestgo/services"
	"github.com/stretchr/testify/assert"
)

func setup() {
	// Setup the test
	productOne := models.Product{
		Name:     "Coffee 1",
		Category: models.Coffee,
		Price:    19,
	}
	productTwo := models.Product{
		Name:     "Coffee 2",
		Category: models.Coffee,
		Price:    29,
	}
	productThree := models.Product{
		Name:     "Equipment 1",
		Category: models.Equipment,
		Price:    79,
	}
	productFour := models.Product{
		Name:     "Accesories 1",
		Category: models.Accesories,
		Price:    59,
	}

	// InsertProduct(product models.Product)
	services.InsertProduct(productOne)
	services.InsertProduct(productTwo)
	services.InsertProduct(productThree)
	services.InsertProduct(productFour)
}

// Case of use 1: use the discount of 10% for the total of the order
// Case of use 2: use the discount of 1 coffee for 2 coffee
// Case of use 3: use the free shipping for 3 equipment
func Test_complete(t *testing.T) {
	// Test the main function
	// Setup the test
	setup()

	// UpsertProduct(product_id int, user_id int, quantity int)
	services.UpsetProduct(1, 1, 1)
	services.UpsetProduct(2, 1, 5)
	services.UpsetProduct(3, 1, 3)
	services.UpsetProduct(4, 1, 2)

	cart, err := services.GetCart(1)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.Equal(t, 4, len(cart.Product_cart), "The cart must have 4 products")

	order, err := services.GetOrder(1)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.Equal(t, float32(519), order.Total.Product, "The total of the product must be 519")
	assert.Equal(t, float32(80.9), order.Total.Discount, "The discount must be 80.9")
	assert.Equal(t, float32(0), order.Total.Shipping, "The shipping must be 0")
	assert.Equal(t, float32(438.1), order.Total.Order, "The total must be 438.1")
}

// Case of use 4: not use any discount
// Case of use 5: not use the free shipping
func Test_without_discount(t *testing.T) {
	// Test the main function without discount

	// Setup the test
	setup()

	// UpsertProduct(product_id int, user_id int, quantity int)
	services.UpsetProduct(1, 1, 1)
	services.UpsetProduct(2, 1, 1)
	services.UpsetProduct(3, 1, 2)
	services.UpsetProduct(4, 1, 1)

	cart, err := services.GetCart(1)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.Equal(t, 4, len(cart.Product_cart), "The cart must have 4 products")

	order, err := services.GetOrder(1)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.Equal(t, float32(265), order.Total.Product, "The total of the product must be 265")
	assert.Equal(t, float32(0), order.Total.Discount, "The discount must be 0")
	assert.Equal(t, float32(5), order.Total.Shipping, "The shipping must be 0")
	assert.Equal(t, float32(270), order.Total.Order, "The total must be 270")
}
