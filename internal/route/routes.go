package route

import "github.com/gofiber/fiber/v2"

func PublicAPI(app *fiber.App) {

}

func PrivateAPI(app *fiber.App) {
	UserRoute(app)
}
