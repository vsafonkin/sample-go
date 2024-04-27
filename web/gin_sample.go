package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunGinServer(host string, port string) {
	r := gin.New()

	r.GET("/", homeGin)
	r.Run(fmt.Sprintf("%s:%s", host, port))
}
func homeGin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello, bob",
	})
}
