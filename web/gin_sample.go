package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RunGinServer() {
	r := gin.Default()

	r.GET("/", homeGin)
	r.Run()
}
func homeGin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello, bob",
	})
}
