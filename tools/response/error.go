package response

import (
	"errors"
)

// Can be added as many as need like belows examples
// 400	BAD_CONTINUATION_TOKEN	Invalid continuation token passed.
// 400	BAD_PAGE	Page number does not exist or is an invalid format (e.g. negative).
// 400	BAD_REQUEST	The resource you’re creating already exists.
// 400	INVALID_ARGUMENT	Invalid argument value passed.
// 400	INVALID_AUTH	Authentication/OAuth token is invalid.
// 400	INVALID_AUTH_HEADER	Authentication header is invalid.
// 400	INVALID_BATCH	Batched request is missing or invalid.
// 400	INVALID_BODY	A request body that was not in JSON format was passed.
// 400	UNSUPPORTED_OPERATION	Requested operation not supported.
// 401	ACCESS_DENIED	Authentication unsuccessful.
// 401	NO_AUTH	Authentication not provided.
// 403	NOT_AUTHORIZED	User has not been authorized to perform that action.
// 404	NOT_FOUND	Invalid URL.
// 405	METHOD_NOT_ALLOWED	Method is not allowed for this endpoint.
// 409	REQUEST_CONFLICT	Requested operation resulted in conflict.
// 429	HIT_RATE_LIMIT	Hourly rate limit has been reached for this token. Default rate limits are 2,000 calls per hour.
// 500	EXPANSION_FAILED	Unhandled error occurred during expansion; the request is likely to succeed if you don’t ask for expansions, but contact Eventbrite support if this problem persists.
// 500	INTERNAL_ERROR	Unhandled error occurred in Eventbrite. contact Eventbrite support if this problem persists.

var (
	ErrNotFound = errors.New("not found")

	ErrInternalServer = errors.New("internal server error")

	ErrAlreadyExists = errors.New(" already exists")

	ErrUsernameExists = errors.New("username exists")

	ErrPhoneExists = errors.New("phone exists")

	ErrEmailExists = errors.New("email exists")

	ErrInvalidField = errors.New("invalid field for username/email")

	ErrMaximumAmount = errors.New("maximum amount")

	ErrNotEnoughCash = errors.New("not enough cash")

	ErrInvalidFieldForOperations = errors.New("invalid field for operation type")

	ErrNotValidPhone = errors.New("invalid field for phone type")

	ErrNotValidFirstName = errors.New("invalid field for firstname type")

	ErrNotValidLastName = errors.New("invalid field for lastname type")

	ErrorNotValidPassword = errors.New("invalid field for Password")

	ErrorSignInCorrect = errors.New("username or password is incorrect")

	ErrorNotANumberLimit = errors.New("query Limit not a number")

	ErrorNotANumberOffset = errors.New("query Offset not a number")

	ErrorNotANumberPage = errors.New("query Page not a number")

	ErrorNotANumberPageSize = errors.New("query PageSize not a number")

	ErrorOffsetNotAUnsignedInt = errors.New("query Offset not a unsigned int")

	ErrorLimitNotAUnsignedInt = errors.New("query Limit not a unsigned int")

	ErrorParamIsEmpty = errors.New("query Parameter is empty")

	ErrorPostIDInvalid = errors.New("post id is invalid")

	ErrorUserIDInvalid = errors.New("user id is invalid")

	ErrorRoleIDInvalid = errors.New("role id is invalid")

	ErrorBranchIDInvalid = errors.New("branch id is invalid")
)
