package schemas

// Password schema structure
type Password struct {
	Hash   string `json:"hash"`
	UserId string `json:"userId" bson:"userId"`
	CommonFields
}
