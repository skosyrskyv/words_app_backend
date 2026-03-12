package dto

type AuthInput struct {
	Email    string
	Password string
}

type RefreshTokenInput struct {
	Access  string
	Refresh string
}
