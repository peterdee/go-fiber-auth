package user

type SignInUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
