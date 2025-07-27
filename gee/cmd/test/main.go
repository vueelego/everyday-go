package main

import (
	"everyday-go/gee"
	"log"
	"net/http"
	"time"
)

func main() {
	engine := gee.New()

	engine.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hi~</h1>")
	})

	engine.POST("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	engine.GET("/json", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"title":   "Hello Gee～",
			"content": "this day2",
			"date":    time.Now().Format("2006/01/02 15:04:05"),
		})
	})

	if err := engine.Run(":4000"); err != nil {
		log.Fatal(err)
	}
}
