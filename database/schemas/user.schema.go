package schemas

// User schema structure
type User struct {
	AvatarLink string `json:"avatarLink"`
	Created    int64  `json:"created"`
	Email      string `json:"email"`
	ID         string `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	Updated    int64  `json:"updated"`
}
