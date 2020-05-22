package request

import (
	"github.com/dgrijalva/jwt-go"
)

// Custom claims structure
type CustomClaims struct {
	UUId     int64
	UserId   int64
	NickName string
	Phone    string
	jwt.StandardClaims
}
