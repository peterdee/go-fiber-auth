package configuration

type ResponseMessagesStruct struct {
	AccessDenied        string
	EmailAlreadyInUse   string
	InternalServerError string
	InvalidData         string
	MissingData         string
	NotFound            string
	Ok                  string
	TooManyRequests     string
}

type RolesStruct struct {
	Admin string
	User  string
}
