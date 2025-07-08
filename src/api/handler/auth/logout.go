package auth

import (
	"net/http"

	"reservation-manager/db/generated"

)

func LogOutHandler(db *generated.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"ログアウト"}`))
	}
}