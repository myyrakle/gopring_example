package dto

type SignupRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
