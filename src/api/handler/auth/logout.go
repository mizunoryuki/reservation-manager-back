package auth

import (
	"net/http"
	"time"

	"reservation-manager/db/generated"
)

//トークンの無効化
func LogOutHandler(db *generated.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			Secure:   true, 
			SameSite: http.SameSiteLaxMode,
			Expires:  time.Unix(0, 0),
			MaxAge:   -1,
		})

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"logged out"}`))
	}
}