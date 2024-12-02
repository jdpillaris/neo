package main

import (
	"net/http"

	"github.com/jdpillaris/neo/handlers"
)

func main() {
	http.HandleFunc("/echo", handlers.Echo)
	http.ListenAndServe(":8080", nil)
}
