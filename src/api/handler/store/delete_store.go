package store

import (
	"net/http"
	"reservation-manager/db/generated"
	"strconv"
	"strings"
)

func DeleteStoreHandler(db *generated.Queries) http.HandlerFunc {
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

		//店舗探索
		_,err = db.GetStoreByID(r.Context(), int32(store_id_int))
		if err != nil {
			http.Error(w,"指定されたidの店舗が見つかりませんでした",http.StatusInternalServerError)
			return
		}

		// 店舗削除
		err = db.DeleteStore(r.Context(),int32(store_id_int))
		if err != nil {
			http.Error(w,"店舗削除に失敗しました", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"正常に店舗を削除することができました"}`))
	}
}
