package dto

import (
	//...
	// import the jwt-go library
	//...
)
// Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
