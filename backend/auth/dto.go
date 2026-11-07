package auth

type RegisterRequest struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
