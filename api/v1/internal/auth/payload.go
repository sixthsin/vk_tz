package auth

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=4,max=16,alphanum"`
	Password string `json:"password" binding:"required,min=8,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
