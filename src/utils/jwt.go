package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"reservation-manager/db/generated"
)

// Claims に含める内容
type Claims struct {
	ID string `json:"user_id"`
	Role string `json:"user_role"`
	jwt.RegisteredClaims
}

// JWTの生成関数
func GenerateJWT(id string,role generated.UsersRole) (string, error) {
	jwtKey := []byte(os.Getenv("JWT_KEY"))

	expirationTime := time.Now().Add(24 * time.Hour)

	
	claims := &Claims{
		ID: id,
		Role: string(role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "reservation-manager",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}
