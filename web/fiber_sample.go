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

func RunFiberServer(host string, port string) {
	r := fiber.New()
	r.Use(logger.New())

	r.Get("/", homeGet)
	r.Get("/:tablename", readData)

	r.Listen(fmt.Sprintf("%s:%s", host, port))
}

func homeGet(c *fiber.Ctx) error {
	return c.JSON(responce{Message: "Hello"})
}

func readData(c *fiber.Ctx) error {
	tablename := c.Params("tablename")
	if tablename == "" {
		return c.JSON(responce{})
	}
	col := c.Query("col")
	val := c.Query("val")
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = %s", tablename, col, val)
	return c.JSON(responce{Message: "Query", Data: query})
}
