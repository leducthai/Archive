package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

func GenerateJWT(UserID string) (string, error) {
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

func verifyJWT(ctx context.Context) (bool, error) {
	if ctx.Value("Token") != nil {
		token, err := jwt.Parse(ctx.Value("Token").(string), func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodECDSA)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return sampleSecretKey, nil
		})
		if err != nil {
			return false, err
		}

		if !token.Valid {
			return false, err
		}
	}
	return false, fmt.Errorf("you're unauthorized")
}
