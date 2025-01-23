package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunGinServer(host string, port string) error {
	r := gin.New()
	r.Use(gin.Logger())

	r.GET("/", homeGin)
	return r.Run(fmt.Sprintf("%s:%s", host, port))
}

func homeGin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello, bob",
	})
}
