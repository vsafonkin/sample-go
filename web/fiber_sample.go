package web

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/logger"
)

type responce struct {
	Message string
}

func RunFiberServer(host string, port string) {
	r := fiber.New()
	// r.Use(logger.New())

	r.Get("/", home)

	r.Listen(fmt.Sprintf("%s:%s", host, port))
}

func home(c *fiber.Ctx) error {
	return c.JSON(responce{Message: "Hello, Bob"})
}
