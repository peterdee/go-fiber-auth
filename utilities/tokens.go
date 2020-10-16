package utilities

import (
	"errors"
	"os"
	"time"

	JWT "github.com/dgrijalva/jwt-go"

	"go-fiber-auth/configuration"
)

// Get JWT secret from the environment
func getSecret() string {
	secret := os.Getenv("TOKENS_ACCESS_SECRET")
	if secret == "" {
		secret = "super-secret"
	}
	return secret
}

// Generate a new JWT token
func GenerateJWT(params GenerateJWTParams) (string, error) {
	expiration := params.ExpiresIn * 60 * 60
	if expiration == 0 {
		expiration = 24 * 60 * 60
	}

	// create claims
	claims := JWTClaims{
		params.UserId,
		JWT.StandardClaims{
			ExpiresAt: time.Now().Unix() + expiration,
		},
	}

	token := JWT.NewWithClaims(JWT.SigningMethodHS256, claims)
	secret := getSecret()

	signedToken, signingError := token.SignedString([]byte(secret))
	if signingError != nil {
		return "", signingError
	}

	return signedToken, nil
}

// Parse JWT, validate it and return claims
func ParseClaims(token string) (*JWTClaims, error) {
	// validate the token and check expiration
	decoded, parsingError := JWT.ParseWithClaims(
		token,
		&JWTClaims{},
		func(decoded *JWT.Token) (interface{}, error) {
			return []byte(getSecret()), nil
		},
	)
	if parsingError != nil {
		return &JWTClaims{}, parsingError
	}

	if claims, ok := decoded.Claims.(*JWTClaims); ok && decoded.Valid {
		return claims, nil
	}
	return &JWTClaims{}, errors.New(configuration.ResponseMessages.InvalidToken)
}
