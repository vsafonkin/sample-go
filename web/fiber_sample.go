package web

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type responce struct {
	Message string
	Data    any
}

type User struct {
	Id   int
	Name string
}

func RunFiberServer(host string, port string) {
	r := fiber.New()
	r.Use(logger.New())

	r.Get("/", homeGet)
	r.Post("/", homePost)

	r.Listen(fmt.Sprintf("%s:%s", host, port))
}

func homeGet(c *fiber.Ctx) error {
	return c.JSON(responce{Message: "Hello"})
}

func homePost(c *fiber.Ctx) error {

	var user User
	if err := c.BodyParser(&user); err != nil {
		fmt.Println("parser error:", err)
	}
	return c.JSON(responce{Message: user.Name, Data: user.Id})
}
