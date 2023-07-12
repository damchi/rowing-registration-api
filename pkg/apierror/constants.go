package apierror

const (
	MsgRouteAccessNotAllowed = "routeAccessNotAllowed"

	MsgValidationFailed   = "validationFailed"
	MsgInsertionFailed    = "insertionFailed"
	MsgTokenRefNotfound   = "tokenRefNotFound"
	MsgEmailAlreadyUsed   = "emailAlreadyUsed"
	MsgValidJsonBody      = "invalidJsonBody"
	MsgCredentialNotValid = "credentialNotValid"

	CodeRouteAccessNotAllowed = 403000
	CodeValidationFailed      = 400001
	CodeNotFound              = 404001
	CodeTokenIsMissing        = 401001

	CodeTokenIsInvalid = 401002
	MsgTokenIsInvalid  = "tokenIsNotValid"
	MsgTokenIsMissing  = "tokenIsMissing"
)
