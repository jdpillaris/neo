package main

import (
	"fmt"
	"net/http"

	"github.com/jdpillaris/neo/handlers"
)

func main() {
	http.HandleFunc("/echo", handlers.EchoMatrix)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
