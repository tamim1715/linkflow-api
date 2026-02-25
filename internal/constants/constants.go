package constants

import "time"

const (
	AppName = "LinkFlow"

	MagicLinkTokenCollection = "magic_link_tokens"
	FeedbackCollection       = "feedbacks"
	UserCollection           = "users"

	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer "

	ContextUserID = "userID"

	ErrInvalidToken = "invalid token"
	ErrMissingToken = "missing token"
	InvalidRequest  = "invalid request"

	APIPrefix         = "/api"
	AuthPrefix        = "/auth"
	VerifyPrefix      = "/verify"
	FeedbackPrefix    = "/feedback"
	RequestLinkPrefix = "/request-link"

	Token = "token"

	AuthRateLimitRequests = 5
	AuthRateLimitWindow   = time.Minute // minutes

	Error     = "error"
	SecretKey = "super-secret-key"
)
