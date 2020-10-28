package schemas

// Password schema structure
type Password struct {
	Hash   string `json:"hash"`
	UserId string `json:"userId" bson:"userId"`
	Created int64  `json:"created"`
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Updated int64  `json:"updated"`
}
