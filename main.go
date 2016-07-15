/**
 * Launch a basic webserver
 * with handling the routes (endpoints)
 * We use here the "mux router"" from the Gorilla Web Toolkit
 */
package main

import (
	"log"
	"net/http"
)

func main() {
	// get the router that handle the appropriate route
	router := GetRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
