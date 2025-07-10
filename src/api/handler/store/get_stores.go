package store

import (
	"encoding/json"
	"log"
	"net/http"
	"reservation-manager/db/generated"
)

func GetStoresHandler(db *generated.Queries) http.HandlerFunc{
return func(w http.ResponseWriter, r *http.Request) {
	stores, err := db.GetAllStores(r.Context())
	if err != nil {
		log.Printf("GetAllStores error: %v",err)
		http.Error(w,"店舗の一覧取得に失敗しました",http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	if err := json.NewEncoder(w).Encode(stores); err != nil {
		http.Error(w, "JSONエンコードに失敗しました", http.StatusInternalServerError)
		return
	}
	}
}