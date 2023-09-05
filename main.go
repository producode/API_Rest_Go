package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"github.com/apirestgo/controller"

	"fmt"
	"os"
)

var app *fiber.App

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	app = fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("PORT_CORS"),
		AllowHeaders: "*",
	}))

	//User
	app.Post("/LogIn", controller.LogIn)

	//Product
	app.Post("/InsertProduct", controller.InsertProduct)
	app.Post("/UpsetProduct", controller.UpsetProduct)
	app.Get("/GetCart", controller.GetCart)
	app.Get("/GetOrder", controller.GetOrder)

	//----Connection test----
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Conexion exitosa")
	})

	err = app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Server running on port " + os.Getenv("PORT"))
}
