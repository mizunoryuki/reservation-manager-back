package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"reservation-manager/db/generated"
	"reservation-manager/models"
)

//新規登録
func SignUpHandler(db *generated.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// JSONデコード
		var input models.SignupInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "無効な入力です", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// 必須チェック
		if input.Email == "" || input.Name == "" || input.Password == "" {
			http.Error(w, "全ての項目を入力してください", http.StatusBadRequest)
			return
		}

		// 重複確認
		_, err := db.GetUserByEmail(r.Context(), input.Email)
		if err != nil && err != sql.ErrNoRows {
			http.Error(w, "サーバーエラー", http.StatusInternalServerError)
			return
		}
		if err == nil {
			http.Error(w, "既に登録されたメールです", http.StatusConflict)
			return
		}

		// パスワードをハッシュ化
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "ハッシュ生成に失敗しました", http.StatusInternalServerError)
			return
		}

		// ユーザー作成
		//sqlcで生成したデータ構造に変換する
		err = db.CreateUser(r.Context(), generated.CreateUserParams{
			Email:         strings.ToLower(input.Email),
			Name:          input.Name,
			PasswordHash:  string(hashedPassword),
		})
		if err != nil {
			http.Error(w, "ユーザー登録に失敗しました", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"登録完了しました"}`))
	}
}
