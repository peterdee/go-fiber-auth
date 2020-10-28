package schemas

// User schema structure
type User struct {
	AvatarLink string `json:"avatarLink"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	Created    int64  `json:"created"`
	ID         string `json:"id,omitempty" bson:"_id,omitempty"`
	Updated    int64  `json:"updated"`
}
