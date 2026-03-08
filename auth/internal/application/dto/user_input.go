package dto

type CrateUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type GetUserUUIDInput struct {
	UUID string `json:"uuid" binding:"required"`
}

type GetUserEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type DeleteUserInput struct {
	UUID string `json:"uuid" binding:"required"`
}

type ChangePasswordInput struct {
	UUID        string `json:"uuid" binding:"required"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
