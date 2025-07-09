package reservation

import (
	"encoding/json"
	"net/http"
	"reservation-manager/db/generated"
	"reservation-manager/middleware"
	"reservation-manager/models"
	"time"
)


func CreateReservationHandler(db *generated.Queries) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		//headerからユーザid取得
		userID,ok := middleware.UserIDFromContext(r.Context())
		if !ok {
			http.Error(w,"UserIDが取得できません",http.StatusUnauthorized)
			return
		}

		// JSONデコード
		var input models.ReservationInfo
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "無効な入力です", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// 必須チェック
		if input.StoreID == 0 || input.VisitDate == "" || input.VisitTime == "" {
			http.Error(w, "全ての項目を入力してください", http.StatusBadRequest)
			return
		}

		//店舗idが正しいか
		_,err := db.GetStoreByID(r.Context(),int32(input.StoreID))
		if err != nil {
			http.Error(w,"指定された店舗idが存在しません",http.StatusBadRequest)
			return
		}

		//ユーザidが正しいか
		_,err = db.GetUserByID(r.Context(),int32(userID)) 
		if err != nil {
			http.Error(w,"指定されたユーザidが存在しません",http.StatusBadRequest)
			return
		}

		visitTime,err := time.Parse("15:04",input.VisitTime)
		if err != nil {
			http.Error(w,"予約時間の形式を変換できませんでした",http.StatusInternalServerError)
			return
		}

		visitTimeStr := visitTime.Format("15:04")

		//dbに格納する方式に予約情報を加工
		datetimeStr := input.VisitDate + "-" + visitTimeStr
		visitDateTime,err := time.Parse("2025-07-08-15:04",datetimeStr)
		if err != nil {
			http.Error(w, "日付フォーマットが不正です", http.StatusBadRequest)
			return
		}

		// 予約作成
		//sqlcで生成したデータ構造に変換する
		err  = db.CreateReservation(r.Context(),generated.CreateReservationParams{
			UserID: int32(userID),
			StoreID: int32(input.StoreID),
			VisitDate: visitDateTime,

		})
		if err != nil {
			http.Error(w,"予約に失敗しました",http.StatusInternalServerError)
			return
		}


		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"正常に予約登録ができました"}`))
	}
}