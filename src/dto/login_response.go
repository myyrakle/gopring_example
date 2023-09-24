package dto

type LoginResponse struct {
	Success     bool   `json:"success"`
	AccessToken string `json:"access_token"`
}
