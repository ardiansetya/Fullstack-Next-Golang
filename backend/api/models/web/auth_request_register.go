package web

type AuthRequestRegister struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
