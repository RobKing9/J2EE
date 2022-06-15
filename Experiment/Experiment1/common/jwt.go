package common

import (
	"Experiment1/models"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

var identitykey = []byte("a_secrect_crect")

type Claims struct {
	Name string
	jwt.StandardClaims
}

func ReleaseToken(u models.User) (string, error) {
	expirationtime := time.Now().Add(100 * 24 * time.Hour)
	claims := &Claims{
		Name: u.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationtime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "oceanlearn.tech",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString(identitykey)
	if err != nil {
		log.Println("err = ", err.Error())
	}
	return tokenstring, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return identitykey, nil
	})

	return token, claims, err
}
