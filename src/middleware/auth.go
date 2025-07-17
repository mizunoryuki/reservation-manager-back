package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"reservation-manager/utils"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "userID"
const UserRoleKey contextKey = "userRole"

//JWTによってルートの制御を行う
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "認証トークンが必要です", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "認証トークンの形式が不正です", http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]
		jwtKey := []byte(os.Getenv("JWT_KEY"))

		// トークンのパースと検証
		claims := &utils.Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "無効なトークンです", http.StatusUnauthorized)
			return
		}

		// コンテキストにユーザ情報を埋め込む
		ctx := context.WithValue(r.Context(), UserIDKey, claims.ID)
		ctx = context.WithValue(ctx, UserRoleKey, claims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserIDFromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(UserIDKey).(string)
	return id, ok
}