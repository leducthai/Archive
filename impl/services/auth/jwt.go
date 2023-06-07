package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

func generateJWT(UserID string) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)

	// comment
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["ID"] = UserID

	// stringify jwt token
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
