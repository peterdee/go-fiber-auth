package utilities

import (
	"os"
	"time"

	JWT "github.com/dgrijalva/jwt-go"
)

// Generate a new JWT token
func GenerateJWT(params GenerateJWTParams) (string, error) {
	token := JWT.New(JWT.SigningMethodHS256)
	claims := token.Claims.(JWT.MapClaims)
	claims["id"] = params.UserId

	expiration := params.ExpiresIn * 60 * 60
	if expiration == 0 {
		expiration = 24 * 60 * 60
	}
	now := time.Now().Unix()
	claims["exp"] = now + expiration

	secret := os.Getenv("TOKENS_ACCESS_SECRET")
	if secret == "" {
		secret = "super-secret"
	}

	signedToken, signingError := token.SignedString([]byte(secret))
	if signingError != nil {
		return "", signingError
	}

	return signedToken, nil
}
