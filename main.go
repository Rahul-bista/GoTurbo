package main

import (
	"GoTurbo/goturbo"
	"net/http"
)

func main() {
	server := goturbo.NewServer()

	// Example middleware
	server.Use(goturbo.LoggingMiddleware)

	server.Handle("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		ctx := goturbo.Context{Writer: w, Request: r}
		ctx.String(http.StatusOK, "Hello, World!")
	})

	server.Handle("GET", "/json", func(w http.ResponseWriter, r *http.Request) {
		ctx := goturbo.Context{Writer: w, Request: r}
		data := map[string]string{"message": "Hello, JSON"}
		ctx.JSON(http.StatusOK, data)
	})

	server.Run(":8080")
}
