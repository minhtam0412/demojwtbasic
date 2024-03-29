package dto
import (
	//...
	// import the jwt-go library
	"github.com/dgrijalva/jwt-go"
	//...
)
// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
