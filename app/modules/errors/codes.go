package errors

const (
	// User Authentication Errors
	CodeInvalidEmail            = 1000
	CodeIncorrectPassword       = 1001
	CodeWeakPassword            = 1002
	CodeUserNotFound            = 1003
	CodeInvalidUsername         = 1004
	CodeUnidenticalPasswordPair = 1005

	// Media Validation Errors
	CodeInvalidMediaType = 2000
	CodeInvalidMediaData = 2001
	CodeMediaTooLarge    = 2002

	// Client Request
	CodeInvalidRequestFormat = 3000
	CodeInvalidOrMissingData = 3001
	CodeMethodNotAllowed     = 3002
	CodeUnauthorized         = 3003

	// Server Errors
	CodeInternalServerError = 4000
	CodeParsingError        = 4001
	CodeUserCreationError   = 4002
)
