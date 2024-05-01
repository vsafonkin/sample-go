package web

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vsafonkin/sample-go/db"
)

type responce struct {
	Message string
	Data    any
}

func RunFiberServer(host string, port string) {
	r := fiber.New()
	r.Use(logger.New())

	conn, err := db.NewConnect(db.Config{
		Host:    "localhost",
		Port:    "5432",
		User:    "postgres",
		Pass:    "admin",
		DBName:  "demo",
		AppName: "goapp",
	})
	if err != nil {
		fmt.Println("db connect error:", err)
	}

	r.Use("/db", func(c *fiber.Ctx) error {
		c.Locals("conn", conn)
		return c.Next()
	})
	r.Get("/", homeGet)

	r.Get("/db/:tablename", getRow)

	r.Listen(fmt.Sprintf("%s:%s", host, port))
}

func homeGet(c *fiber.Ctx) error {
	return c.JSON(responce{Message: "Hello"})
}

func getRow(c *fiber.Ctx) error {
	conn := c.Locals("conn").(*db.DB)
	tablename := c.Params("tablename")
	if tablename == "" {
		return c.JSON(responce{})
	}
	col := c.Query("col")
	val := c.Query("val")
	row := conn.Row(tablename, col, val)
	return c.JSON(responce{Message: "Query", Data: row})
}
