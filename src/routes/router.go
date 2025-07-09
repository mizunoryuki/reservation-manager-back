package routes

import (
	"net/http"

	"reservation-manager/api/handler/auth"
	"reservation-manager/api/handler/reservation"
	"reservation-manager/api/handler/store"
	"reservation-manager/db/generated"
	"reservation-manager/middleware"
)

func InitRoutes(q *generated.Queries) *http.ServeMux {
	mux := http.NewServeMux()

	// ユーザ関連
	mux.HandleFunc("/signup", auth.SignUpHandler(q))
	mux.HandleFunc("/login", auth.LogInHandler(q))
	mux.HandleFunc("/logout", auth.LogOutHandler(q))

	// 店舗情報
	mux.Handle("/admin/stores", middleware.AuthMiddleware(
		middleware.RoleMiddleware("admin")(http.HandlerFunc(store.CreateStoreHandler(q))),
	))
	mux.Handle("/admin/stores/", middleware.AuthMiddleware(
		middleware.RoleMiddleware("admin")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPut:
				store.UpdateStoreHandler(q)(w, r)
			case http.MethodDelete:
				store.DeleteStoreHandler(q)(w, r)
			default:
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
		})),
	))
	mux.Handle("/stores", middleware.AuthMiddleware(
		middleware.RoleMiddleware("admin", "general")(http.HandlerFunc(store.GetStoresHandler(q))),
	))

	// 予約情報（管理者）
	mux.Handle("/admin/reservations", middleware.AuthMiddleware(
		middleware.RoleMiddleware("admin")(http.HandlerFunc(reservation.GetReservationsHandler(q))),
	))
	mux.Handle("/admin/reservations/", middleware.AuthMiddleware(
		middleware.RoleMiddleware("admin")(http.HandlerFunc(reservation.DeleteReservationHandler(q))),
	))

	// 予約情報（一般ユーザ）
	mux.Handle("/user/reservations", middleware.AuthMiddleware(
		middleware.RoleMiddleware("general")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost:
				reservation.CreateReservationHandler(q)(w, r)
			case http.MethodGet:
				reservation.GenGetReservationsHandler(q)(w, r)
			default:
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
		})),
	))
	mux.Handle("/user/reservations/", middleware.AuthMiddleware(
		middleware.RoleMiddleware("general")(http.HandlerFunc(reservation.GenDeleteReservationHandler(q))),
	))

	return mux
}
