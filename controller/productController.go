package controller

import (
	"github.com/apirestgo/models"
	"github.com/apirestgo/services"
	"github.com/gofiber/fiber/v2"

	"strconv"
)

type Product_cart_dto struct {
	Quantity int
}

type Product_dto struct {
	Name     string
	Category string
	Price    float32
}

func InsertProduct(c *fiber.Ctx) error {
	product_dto := new(Product_dto)
	if err := c.BodyParser(product_dto); err != nil {
		return err
	}
	product := models.Product{
		Name:     product_dto.Name,
		Category: product_dto.Category,
		Price:    product_dto.Price,
	}
	err, product := services.InsertProduct(product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.JSON(product)
}

func UpsertProductCart(c *fiber.Ctx) error {
	user_id, err := services.Authenticate(c.Get("Authorization"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	product_id, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	err, cart := services.AddToCart(product_id, user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.JSON(cart)
}

func ModifyCart(c *fiber.Ctx) error {
	product_cart_dto := new(Product_cart_dto)
	if err := c.BodyParser(product_cart_dto); err != nil {
		return err
	}
	user_id, err := services.Authenticate(c.Get("Authorization"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	product_id, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	err, cart := services.ModifyCartQuantity(product_id, user_id, product_cart_dto.Quantity)
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
