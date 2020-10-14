package auth

type SignInUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpUserRequest struct {
	Name string `json:"name"`
	Role string `json:"role"`
	SignInUserRequest
}
