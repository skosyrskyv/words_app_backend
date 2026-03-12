package dto

type CrateUserInput struct {
	Email    string
	Password string
}

type GetUserUUIDInput struct {
	UUID string
}

type GetUserEmailInput struct {
	Email string
}

type DeleteUserInput struct {
	UUID string
}

type ChangePasswordInput struct {
	UUID        string
	OldPassword string
	NewPassword string
}
