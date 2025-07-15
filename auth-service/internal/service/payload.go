package service

type UserResponse struct {
	Username string `json:"username"`
	IsValid  bool   `json:"is_valid"`
}
