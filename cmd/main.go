package main

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"wms-api-v2/internal/route"
	"wms-api-v2/pkg/database"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
	})

	database.MysqlConn()
	route.PrivateAPI(app)
	route.PublicAPI(app)

	err := app.Listen("localhost:8000")
	if err != nil {
		panic("app start error")
	}
}
