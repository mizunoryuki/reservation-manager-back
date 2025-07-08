package store

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"reservation-manager/db/generated"
)

// UpdateStoreHandler 店舗情報更新ハンドラ
func UpdateStoreHandler(db *generated.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ルートパラメータからstore_idを取得
		path := r.URL.Path
		parts := strings.Split(path,"/")// `/admin/stores/:id`

		store_id_str := parts[len(parts)-1]

		store_id_int, err := strconv.Atoi(store_id_str)
		if err != nil {
			http.Error(w, "store_idが不正です", http.StatusBadRequest)
			return
		}

		store_id_i32 := int32(store_id_int)

		//店舗検索
		_,err = db.GetStoreByID(r.Context(),store_id_i32)
		if err != nil {
			http.Error(w,"店舗が見つかりませんでした",http.StatusBadRequest)
			return
		}

		// リクエストボディから更新データを取得
		var input struct {
			Name              string `json:"name"`
			Address           string `json:"address"`
			BusinessStartTime string `json:"business_start_time"`
			BusinessEndTime   string `json:"business_end_time"`
			Details           string `json:"details"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "無効な入力です", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		//時間文字列を事項形式に変換する
		startTime, err := time.Parse("15:04", input.BusinessStartTime)
		if err != nil {
			http.Error(w, "開始時間の形式が不正です", http.StatusBadRequest)
			return
		}
		endTime, err := time.Parse("15:04", input.BusinessEndTime)
		if err != nil {
			http.Error(w, "終了時間の形式が不正です", http.StatusBadRequest)
			return
		}

		// 更新処理
		err = db.UpdateStore(r.Context(), generated.UpdateStoreParams{
			ID:                store_id_i32,
			Name:              input.Name,
			Address:           input.Address,
			BusinessStartTime: startTime,
			BusinessEndTime:   endTime,
			Details:           sql.NullString{String : input.Details,Valid: input.Details != ""},
		})
		if err != nil {
			http.Error(w, "店舗情報の更新に失敗しました", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"更新が完了しました"}`))
	}
}
