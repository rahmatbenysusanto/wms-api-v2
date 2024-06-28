package main

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"hello": "world",
		})
	})

	err := app.Listen("localhost:8000")
	if err != nil {
		panic("app start error")
	}
}
