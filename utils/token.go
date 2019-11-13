package utils

import (
	"time"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)


func CreateToken(uid uint) (string, error){
	claims := jwt.MapClaims{}
	claims["user_id"] =  uid
	claims["exp"] = time.Now().Add(time.Hour *24).Unix()
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	return token.SignedString([]byte(os.Getenv("SECRET")))
}
