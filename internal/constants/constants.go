package constants

const (
	// App
	AppName = "LinkFlow"

	// Database
	MagicLinkTokenCollection = "magic_link_tokens"
	FeedbackCollection       = "feedbacks"

	// Headers
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer "

	// Context keys
	ContextUserID = "userID"

	// Errors
	ErrInvalidToken = "invalid token"
	ErrMissingToken = "missing token"

	// Routes
	APIPrefix         = "/api"
	AuthPrefix        = "/auth"
	VerifyPrefix      = "/verify"
	FeedbackPrefix    = "/feedback"
	RequestLinkPrefix = "/request-link"

	// Extra
	Token = "token"
)
