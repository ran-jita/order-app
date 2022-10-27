package dto

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
