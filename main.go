package main

import (
	"fmt"
	"net/http"

	"github.com/jdpillaris/neo/handlers"
)

func main() {
	setupRoutes()

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func setupRoutes() {
	http.HandleFunc("POST /echo", handlers.EchoMatrix)
	http.HandleFunc("POST /invert", handlers.TransposeMatrix)
	http.HandleFunc("POST /flatten", handlers.FlattenMatrix)
	http.HandleFunc("POST /add", handlers.AddMatrixElems)
	http.HandleFunc("POST /multiply", handlers.MultiplyMatrixElems)
}
