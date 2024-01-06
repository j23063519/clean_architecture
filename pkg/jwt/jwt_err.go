package jwt

import "errors"

var (
	ErrTokenExpired              error = errors.New("token expired")
	ErrTokenBlacklistGracePeriod error = errors.New("token blacklist grace period")
	ErrTokenExpiredMaxRefresh    error = errors.New("token refresh time expired")
	ErrTokenMalformed            error = errors.New("token malformed")
	ErrTokenInvalid              error = errors.New("token invalid")
	ErrHeaderEmpty               error = errors.New("header empty")
	ErrHeaderMalformed           error = errors.New("header authrization failed")
)
