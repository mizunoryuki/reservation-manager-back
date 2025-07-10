package store

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"reservation-manager/db/generated"
	"reservation-manager/models"
)

//新規登録
func CreateStoreHandler(db *generated.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// JSONデコード
		var input models.StoreInfo
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "無効な入力です", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// 必須チェック
		if input.Name == "" || input.Address == "" || input.BusinessStartTime=="" || input.BusinessEndTime=="" {
			http.Error(w, "全ての項目を入力してください", http.StatusBadRequest)
			return
		}

		//時間文字列を時間に変換
		startTime, err := time.Parse("2006-01-02 15:04", "2025-07-10" + " "+ input.BusinessStartTime)
		if err != nil {
			log.Printf("startTime error: %v", err)
			http.Error(w, "開始時間の形式が不正です", http.StatusBadRequest)
			return
		}
		endTime, err := time.Parse("2006-01-02 15:04", "2025-07-10" + " "+ input.BusinessEndTime)
		if err != nil {
			log.Printf("endTime error: %v", err)
			http.Error(w, "終了時間の形式が不正です", http.StatusBadRequest)
			return
		}

		// 店舗作成
		//sqlcで生成したデータ構造に変換する
		err = db.CreateStore(r.Context(),generated.CreateStoreParams{
			Name: input.Name,
			Address: input.Address,
			BusinessStartTime: startTime,
			BusinessEndTime: endTime,
			Details: sql.NullString{String: input.Details, Valid: input.Details != ""},
		})
		if err != nil {
			log.Printf("CreateStore error: %v", err)
			http.Error(w,"店舗登録に失敗しました", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"正常に店舗の登録ができました"}`))
	}
}
