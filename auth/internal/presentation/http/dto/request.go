package dto

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type GetUserRequest struct {
	UUID string `uri:"uuid" binding:"required"`
}
