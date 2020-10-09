package models

// User schema structure
type User struct {
	Created int64  `json:"created"`
	Email   string `json:"email"`
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Updated int64  `json:"updated"`
}
