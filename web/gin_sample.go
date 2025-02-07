package web

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

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
	cmd := exec.Command("cat", "/proc/cpuinfo")
	stdout, err := cmd.Output()
	if err != nil {
		log.Println(err.Error())
	}
	s := strings.Split(string(stdout), "\n")
	procs := make([]string, 0)
	for _, line := range s {
		if strings.Contains(line, "cpu MHz") {
			lines := strings.Split(line, ":")
			f := strings.Trim(lines[1], " ")
			procs = append(procs, f)
		}
	}
	output := ""
	for i, v := range procs {
		label := fmt.Sprintf("core=\"%d\"", i)
		metric := fmt.Sprintf("localhost_cpu{%s} %s\n", label, v)
		output += metric
	}
	c.Data(200, "text/plain; charset=utf-8", []byte(output))
}
