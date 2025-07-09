package reservation

import (
	"encoding/json"
	"net/http"
	"reservation-manager/db/generated"
)



func GetReservationHandler(db *generated.Queries) http.HandlerFunc{
return func(w http.ResponseWriter, r *http.Request) {
	//予約取得
	reservations, err := db.GetAllReservations(r.Context())
	if err != nil {
		http.Error(w,"予約の一覧取得に失敗しました",http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	if err := json.NewEncoder(w).Encode(reservations); err != nil {
		http.Error(w, "JSONエンコードに失敗しました", http.StatusInternalServerError)
		return
	}
	}
}