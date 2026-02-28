package constants

import "time"

const (
	MagicLinkTokenCollection = "magic_link_tokens"
	FeedbackCollection       = "feedbacks"
	UserCollection           = "users"

	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer "

	ErrInvalidToken     = "invalid token"
	ErrMissingToken     = "missing token"
	InvalidRequest      = "invalid request"
	InvalidEmailAddress = "invalid email address"
	TokenAlreadyUsed    = "token already used"
	TokenExpired        = "token expired"
	TooManyRequests     = "too many requests, please try again later"

	MagicLinkSent     = "magic link sent"
	FeedbackSubmitted = "feedback submitted"

	APIPrefix         = "/api"
	AuthPrefix        = "/auth"
	VerifyPrefix      = "/verify"
	FeedbackPrefix    = "/feedback"
	RequestLinkPrefix = "/request-link"

	AuthRateLimitRequests = 5
	AuthRateLimitWindow   = time.Minute // minutes

	Token         = "token"
	Error         = "error"
	Message       = "message"
	ContextUserID = "userID"
)
