package main

import (
	"app/app/routes"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("-- Starting HTTP Server --")
	fmt.Println("-- Parsing Routes --")

	routes.LoadRoutes()

	fmt.Println("-- Parsed Routes --")
	fmt.Println("-- Serving HTTP Server --")

	fmt.Println("127.0.0.1:8080")
	http.ListenAndServe(":8080", routes.GetRouter())
}
