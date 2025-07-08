package store

import (
	"encoding/json"
	"net/http"
	"reservation-manager/db/generated"
)

func GetStoresHandler(db *generated.Queries) http.HandlerFunc{
return func(w http.ResponseWriter, r *http.Request) {
	stores, err := db.GetAllStores(r.Context())
	if err != nil {
		http.Error(w,"店舗の一覧取得に失敗しました",http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type","application/json")
	if stores != nil {
	json.NewEncoder(w).Encode(stores)
	}else {
		json.NewEncoder(w).Encode(0)
	}
	}
}