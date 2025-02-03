package web

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func RunGinServer(host string, port string) error {
	r := gin.New()
	r.Use(gin.Logger())

	r.GET("/", homeGin)
	r.GET("/exec", execCommand)
	r.GET("/metrics", Metrics)

	return r.Run(fmt.Sprintf("%s:%s", host, port))
}

func homeGin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello, bob",
	})
}

func execCommand(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	queryCommand := queryParams.Get("cmd")
	cmd := exec.Command(queryCommand)
	stdout, err := cmd.Output()
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}
	c.JSON(http.StatusOK, gin.H{
		"command": queryCommand,
		"stdout":  string(stdout),
		"stderr":  errorMessage,
	})
}

func Metrics(c *gin.Context) {
	data := "hello, alice"
	c.Data(200, "text/plain; charset=utf-8", []byte(data))
}
