package schemas

// Password schema structure
type Password struct {
	Created int64  `json:"created"`
	Hash    string `json:"hash"`
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Updated int64  `json:"updated"`
	UserId  string `json:"userId"`
}
