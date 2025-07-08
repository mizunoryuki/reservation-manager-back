package models

//登録する店舗データ
type StoreInfo struct {
	Name     string `json:"name"`
	Address string `json:"address"`
	BusinessStartTime string `json:"business_start_time"`
	BusinessEndTime string `json:"business_end_time"` 
	Details string `json:"details"`
}
