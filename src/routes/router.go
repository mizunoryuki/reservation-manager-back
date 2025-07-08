package routes

import (
	"net/http"
	"strings"

	"reservation-manager/api/handler/auth"
	"reservation-manager/api/handler/store"
	"reservation-manager/db/generated"
	"reservation-manager/middleware"
)

func handleStoreRoutes(mux *http.ServeMux, q *generated.Queries) {
	mux.Handle("/admin/stores", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			store.CreateStoreHandler(q)(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})))

	// /admin/stores/{id} に対する PUT（更新）・DELETE（削除）を処理
	mux.Handle("/admin/stores/", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/admin/stores/")
		if id == "" {
			http.Error(w, "IDが指定されていません", http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodPut:
			store.UpdateStoreHandler(q)(w, r)
		case http.MethodDelete:
			store.DeleteStoreHandler(q)(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})))

	mux.Handle("/stores", middleware.AuthMiddleware(store.GetStoresHandler(q)))
}

func InitRoutes(q *generated.Queries) *http.ServeMux {
	mux := http.NewServeMux()

	// ユーザ関連ルート
	mux.HandleFunc("/signup", auth.SignUpHandler(q))
	mux.HandleFunc("/login", auth.LogInHandler(q))
	mux.HandleFunc("/logout", auth.LogOutHandler(q))

	// 店舗関連ルート
	handleStoreRoutes(mux, q)

	return mux
}
