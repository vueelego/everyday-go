package main

import (
	"everyday-go/gee"
	"fmt"
	"log"
	"net/http"
)

func main() {
	engine := gee.New()

	engine.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi~")
	})

	engine.POST("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})

	if err := engine.Run(":4000"); err != nil {
		log.Fatal(err)
	}
}
