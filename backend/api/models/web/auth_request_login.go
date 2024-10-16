package web

type AuthRequestLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
