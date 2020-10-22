package schemas

// Common schema fields
type CommonFields struct {
	Created int64  `json:"created"`
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Updated int64  `json:"updated"`
}
