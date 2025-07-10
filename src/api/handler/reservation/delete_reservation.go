package reservation

import (
	"log"
	"net/http"
	"reservation-manager/db/generated"
	"reservation-manager/middleware"
	"strconv"
	"strings"
)

func DeleteReservationHandler(db *generated.Queries) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		//ユーザid取得
		userID,ok := middleware.UserIDFromContext(r.Context())
		if !ok {
			http.Error(w,"UserIDが取得できません",http.StatusUnauthorized)
			return
		}

		//パスパラメータから予約id取得
		path := r.URL.Path
		parts := strings.Split(path,"/")// `/admin/reservations/:id`
		res_id_str := parts[len(parts)-1]
		res_id_int, err := strconv.Atoi(res_id_str)
		if err != nil {
			http.Error(w, "予約idが不正です", http.StatusBadRequest)
			return
		}
		//予約idが正しいか
		_,err = db.GetReservationByID(r.Context(),int32(res_id_int))
		if err != nil {
			http.Error(w,"指定された予約idが存在しません",http.StatusBadRequest)
			return
		}


		// 予約削除
		//sqlcで生成したデータ構造に変換する
		err  = db.DeleteReservationAsAdmin(r.Context(),int32(userID))
		if err != nil {
			log.Printf("DeleteReservationAsAdmin error : %v",err)
			http.Error(w,"予約削除に失敗しました",http.StatusInternalServerError)
			return
		}


		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"正常に予約削除ができました"}`))
	}
}

//一般ユーザが予約を削除するための関数
func GenDeleteReservationHandler(db *generated.Queries) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		//ユーザid取得
		userID,ok := middleware.UserIDFromContext(r.Context())
		if !ok {
			http.Error(w,"UserIDが取得できません",http.StatusUnauthorized)
			return
		}

		//パスパラメータから予約id取得
		path := r.URL.Path
		parts := strings.Split(path,"/")// `/user/reservations/:id`
		res_id_str := parts[len(parts)-1]
		res_id_int, err := strconv.Atoi(res_id_str)
		if err != nil {
			http.Error(w, "予約idが不正です", http.StatusBadRequest)
			return
		}

		//予約idが正しいか
		reservation,err := db.GetReservationByID(r.Context(),int32(res_id_int))
		if err != nil {
			http.Error(w,"指定された予約idが存在しません",http.StatusBadRequest)
			return
		}
		//予約の照合
		if reservation.UserID != int32(userID){
			http.Error(w,"ユーザと予約者が異なります",http.StatusForbidden)
			return
		}

		err = db.CancelReservation(r.Context(),generated.CancelReservationParams{
			ID: int32(res_id_int),
			UserID: int32(userID),
		})
		if err != nil {
			log.Printf("CancelReservation error : %v",err)
			http.Error(w,"予約の削除に失敗しました",http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"正常に予約削除ができました"}`))
	}
}