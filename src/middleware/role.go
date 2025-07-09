package middleware

import (
	"net/http"
)

func RoleMiddleware(allowedRoles ...string) func(http.Handler) http.Handler{
	roleSet := make(map[string]struct{})
	for _,role := range allowedRoles {
		roleSet[role] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role,ok := r.Context().Value(UserRoleKey).(string)
			if !ok {
				http.Error(w,"ロール情報がありません",http.StatusUnauthorized)
				return
			}
			
			if _,allowed :=roleSet[role]; !allowed {
				http.Error(w,"許可されていないロールです",http.StatusForbidden)
				return
			}
			next.ServeHTTP(w,r)
		})
	}
}