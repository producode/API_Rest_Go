package controller

import (
	"github.com/apirestgo/models"
	"github.com/apirestgo/services"
	"github.com/gofiber/fiber/v2"
)

func InsertProduct(c *fiber.Ctx) error {
	product_dto := new(models.Product_dto)
	if err := c.BodyParser(product_dto); err != nil {
		return err
	}
	product := &models.Product{
		Name:     product_dto.Name,
		Category: product_dto.Category,
		Price:    product_dto.Price,
	}
	err, id := services.InsertProduct(*product)
	product.Id = id
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.JSON(product)
}

func UpsetProduct(c *fiber.Ctx) error {
	product_cart_dto := new(models.Product_cart_dto)
	if err := c.BodyParser(product_cart_dto); err != nil {
		return err
	}
	user_id, err := services.Authenticate(c.Get("Authorization"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	quantity := product_cart_dto.Quantity
	product_id := product_cart_dto.Id
	err, cart := services.UpsetProduct(product_id, user_id, quantity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.JSON(cart)
}

func GetCart(c *fiber.Ctx) error {
	user_id, err := services.Authenticate(c.Get("Authorization"))
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(err)
	}
	cart, err := services.GetCart(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.JSON(cart)
}

func GetOrder(c *fiber.Ctx) error {
	user_id, err := services.Authenticate(c.Get("Authorization"))
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(err)
	}
	order, err := services.GetOrder(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.JSON(order)
}
