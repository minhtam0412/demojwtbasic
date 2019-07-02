package main

import (
	"MyAPI/controller"
	"log"
	"net/http"
)

//https://www.sohamkamani.com/blog/golang/2019-01-01-jwt-authentication/
func main() {
	// "Signin" and "Welcome" are the handlers that we will implement
	http.HandleFunc("/signin", controller.Signin)
	http.HandleFunc("/welcome", controller.Welcome)
	http.HandleFunc("/refresh", controller.Refresh)

	// start the server on port 4587
	log.Fatal(http.ListenAndServe(":4587", nil))
}
