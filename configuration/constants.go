package configuration

// Server response messages
var ResponseMessages = ResponseMessagesStruct{
	EmailAlreadyInUse: "EMAIL_ALREADY_IN_USE",
	InternalServerError: "INTERNAL_SERVER_ERROR",
	InvalidData:         "INVALID_DATA",
	MissingData:         "MISSING_DATA",
	NotFound:            "NOT_FOUND",
	Ok:                  "OK",
	TooManyRequests:     "TOO_MANY_REQUESTS",
}

// Available user roles
var Roles = RolesStruct{
	Admin: "admin",
	User:  "user",
}
