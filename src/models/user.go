package models

//ユーザ登録時に必要なデータ
type SignupInput struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

//ログイン時に使用する
type LoginInput struct {
	Email	string	`json:"email"`
	Password 	string	`json:"password"`
}
