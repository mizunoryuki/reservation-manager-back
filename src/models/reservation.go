package models

//登録する予約データ
type ReservationInfo struct {
	StoreID int `json:"store_id"`
	VisitDate string `json:"visit_date"`
	VisitTime string `json:"visit_time"` //予約の開始時間(ex.10:00~11:00だったら10:00)
}