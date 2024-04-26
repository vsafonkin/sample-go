package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type responce struct {
	Message string
}

func RunFiberServer() {
	r := fiber.New()
	r.Use(logger.New())

	r.Get("/", home)

	r.Listen(":8080")
}

func home(c *fiber.Ctx) error {
	return c.JSON(responce{Message: "Hello, Bob"})
}
