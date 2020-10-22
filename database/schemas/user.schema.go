package schemas

// User schema structure
type User struct {
	AvatarLink string `json:"avatarLink"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	CommonFields
}
