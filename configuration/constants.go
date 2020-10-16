package configuration

// Server response messages
var ResponseMessages = ResponseMessagesStruct{
	AccessDenied:        "ACCESS_DENIED",
	EmailAlreadyInUse:   "EMAIL_ALREADY_IN_USE",
	InternalServerError: "INTERNAL_SERVER_ERROR",
	InvalidData:         "INVALID_DATA",
	InvalidToken:        "INVALID_TOKEN",
	MissingData:         "MISSING_DATA",
	MissingToken:        "MISSING_TOKEN",
	NotFound:            "NOT_FOUND",
	Ok:                  "OK",
	TooManyRequests:     "TOO_MANY_REQUESTS",
}

// Available user roles
var Roles = RolesStruct{
	Admin: "admin",
	User:  "user",
}
