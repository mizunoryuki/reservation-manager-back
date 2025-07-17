package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"reservation-manager/db/generated"
	"reservation-manager/models"
	"reservation-manager/utils"
)

func LogInHandler(db *generated.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// JSONデコード
		var input models.LoginInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "無効な入力です", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		//必須項目
		if input.Email == "" || input.Password == "" {
			http.Error(w, "全ての項目を入力してください", http.StatusBadRequest)
			return
		}
		
		//ユーザ検索
		user, err := db.GetUserByEmail(r.Context(), input.Email)
		if err != nil {
			if err != sql.ErrNoRows {
				http.Error(w, "ユーザが存在しません", http.StatusInternalServerError)
				return
			}
			http.Error(w,"サーバーエラー",http.StatusInternalServerError)
			return
		}

		//パスワードの照合
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
			http.Error(w, "パスワードが間違っています", http.StatusUnauthorized)
			return
		}


		//JWT発行
		accessToken, err := utils.GenerateJWT(strconv.Itoa(int(user.ID)),user.Role)
		if err != nil {
			http.Error(w,"トークン生成に失敗しました",http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    accessToken,
			Path:     "/",
			HttpOnly: true,
			Secure:   false, 
			SameSite: http.SameSiteLaxMode,
			MaxAge:   60 * 60 * 24,    
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message":      "ログイン完了",
		})
	}
}
