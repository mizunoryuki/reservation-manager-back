package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"reservation-manager/db/generated"
)

// Claims に含める内容
type Claims struct {
	Email string `json:"email"`
	Role string `json:"password"`
	jwt.RegisteredClaims
}

// JWTの生成関数
func GenerateJWT(email string,role generated.UsersRole) (string, error) {
	jwtKey := []byte(os.Getenv("JWT_KEY"))

	expirationTime := time.Now().Add(24 * time.Hour)

	
	claims := &Claims{
		Email: email,
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
