package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/greet/:name", func(c *gin.Context) {
		name := c.Param("name")
		msg, err := greetUser(name)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusOK, msg)
	})
	r.GET("/timestamp", func(c *gin.Context) {
		msg, err := getCurrentTime()
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusOK, msg)
	})
	r.GET("/health", func(c *gin.Context) {
		msg, err := healthStatus()
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusOK, msg)
	})
	r.Run()
}

func greetUser(name string) (string, error) {
	return fmt.Sprintf("Hello, %s", name), nil
}

func getCurrentTime() (string, error) {
	now := time.Now()
	return "The time is " + now.Format(time.UnixDate), nil
}

func healthStatus() (string, error) {
	return "ok", nil
}
