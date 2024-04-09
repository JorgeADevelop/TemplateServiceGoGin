package utils

import (
	"os"

	"github.com/golang-jwt/jwt"
)

func validateJWTToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

func ExtractClaimsFromToken(token string) map[string]interface{} {
	parsedToken, _ := validateJWTToken(token)
	claims := parsedToken.Claims.(jwt.MapClaims)
	return claims
}
