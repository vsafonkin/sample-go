package web

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vsafonkin/sample-go/db"
)

type response struct {
	Message string
	Data    any
}

func RunFiberServer(host string, port string) error {
	r := fiber.New()
	r.Use(logger.New())

	dbname := "demo"
	conn, err := db.NewConnect(db.Config{
		Host:    "localhost",
		Port:    "5432",
		User:    "postgres",
		Pass:    "admin",
		DBName:  dbname,
		AppName: "goapp",
	})
	if err != nil {
		fmt.Println("db connect error:", err)
	}

	r.Get("/", homeGet)

	dbRoute := fmt.Sprintf("/%s", dbname)
	r.Use(dbRoute, func(c *fiber.Ctx) error {
		c.Locals("conn", conn)
		return c.Next()
	})
	tableRoute := fmt.Sprintf("%s/:tablename", dbRoute)
	r.Get(tableRoute, getRow)

	return r.Listen(fmt.Sprintf("%s:%s", host, port))
}

func homeGet(c *fiber.Ctx) error {
	return c.JSON(response{Message: "Hello"})
}

func getRow(c *fiber.Ctx) error {
	conn := c.Locals("conn").(*db.DB)
	tablename := c.Params("tablename")
	if tablename == "" {
		return c.JSON(response{})
	}
	col := c.Query("col")
	val := c.Query("val")
	row := conn.Row(tablename, col, val)
	return c.JSON(response{Message: tablename, Data: row})
}
