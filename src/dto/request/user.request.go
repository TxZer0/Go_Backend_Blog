package request

type Register struct {
	Email          string `json:"email" binding:"required"`
	Password       string `json:"password" binding:"required"`
	VerifyPassword string `json:"verify_password" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
