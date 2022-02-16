package middleware

import (
	"errors"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

// secret key being used to sign tokens
var (
	SecretKey = []byte("buburdepok")
)

// GenerateToken generates a jwt token and assign a username to it's claims and return it
func GenerateToken(userId string) (string, error) {
	/* Create a map to store our claims */
	payload := jwt.MapClaims{}
	/* Set token claims */
	payload["id"] = userId
	payload["exp"] = time.Now().Add(time.Hour * 144).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return signedToken, err
	}
	return signedToken, nil
}

// ParseToken parses a jwt token and returns the username in it's claims
func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return SecretKey, nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}
