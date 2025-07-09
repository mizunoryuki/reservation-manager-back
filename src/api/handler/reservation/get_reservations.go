package reservation

import (
	"encoding/json"
	"net/http"
	"reservation-manager/db/generated"
	"reservation-manager/middleware"
)



func GetReservationsHandler(db *generated.Queries) http.HandlerFunc{
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

//一般ユーザが予約情報を取得する関数
func GenGetReservationsHandler(db *generated.Queries) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		//headerからユーザid取得
		userID,ok := middleware.UserIDFromContext(r.Context())
		if !ok {
			http.Error(w,"UserIDが取得できません",http.StatusUnauthorized)
			return
		}

		reservations,err := db.GetReservationByID(r.Context(),int32(userID))
		if err != nil {
			http.Error(w,"ユーザidから予約情報の取得に失敗しました",http.StatusForbidden)
			return
		}
		w.Header().Set("Content-Type","application/json")
		if err := json.NewEncoder(w).Encode(reservations); err != nil {
			http.Error(w, "JSONエンコードに失敗しました", http.StatusInternalServerError)
			return
		}
	}
}