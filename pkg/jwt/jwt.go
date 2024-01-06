package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/j23063519/clean_architecture/config"
	"github.com/j23063519/clean_architecture/pkg/log"
	"github.com/spf13/cast"
)

// custom JWT struct
type JWT struct {
	Key        []byte
	MaxRefresh time.Duration
}

// JWTCustomClaims custom claims
type JWTCustomClaims struct {
	TableName string    `json:"table_name"`
	UserID    string    `json:"user_id"`
	UserName  string    `json:"user_name"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`

	// the `iss` (Issuer)
	// the `sub` (Subject)
	// the `aud` (Audience)
	// the `exp` (Expiration Time)
	// the `nbf` (Not Before)
	// the `iat` (Issued At)
	// the `jti` (JWT ID)
	jwtpkg.RegisteredClaims
}

// new JWT
func NewJWT() *JWT {
	return &JWT{
		Key:        []byte(config.Config.App.KEY),
		MaxRefresh: time.Duration(cast.ToInt64(config.Config.JWT.MAXREFRESHTIME)) * time.Minute,
	}
}

// new JWTCustomClaims
func newJWTCustomClaims(tablename, userid, username string, now time.Time, expireTime time.Duration) JWTCustomClaims {
	return JWTCustomClaims{
		TableName: tablename,
		UserID:    userid,
		UserName:  username,
		IssuedAt:  now,
		ExpiredAt: now.Add(expireTime),
		RegisteredClaims: jwtpkg.RegisteredClaims{
			Issuer:    config.Config.App.NAME,
			Subject:   "authorization",
			ExpiresAt: jwtpkg.NewNumericDate(now.Add(expireTime)),
			IssuedAt:  jwtpkg.NewNumericDate(now),
			NotBefore: jwtpkg.NewNumericDate(now),
			ID:        uuid.NewString(),
		},
	}
}

// create token for internal
func (jwt *JWT) createToken(claims JWTCustomClaims) (string, error) {
	// use HS256 create token
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwt.Key))
}

// get current time.Time
func (jwt *JWT) getTimeWithLocation() time.Time {
	timezone, _ := time.LoadLocation(config.Config.App.TIMEZONE)
	return time.Now().In(timezone)
}

// create token for outside
func (jwt *JWT) IssueToken(tablename, userid, username string) string {
	claims := newJWTCustomClaims(
		tablename, userid, username,
		jwt.getTimeWithLocation(),
		time.Duration(cast.ToInt64(config.Config.JWT.EXPIRETIME))*time.Minute,
	)
	token, err := jwt.createToken(claims)
	if err != nil {
		log.ErrorJSON("JWT Token", "IssueToken", err)
		return ""
	}
	return token
}

// getTokenFromHeader Using jwtpkg.ParseWithClaims to parse the token from request header
//
// Authorization:Bearer xxxxx
func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}

	// split from blank
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}
	return parts[1], nil
}

// parse token for internal
func (jwt *JWT) parseToken(token string) (*jwtpkg.Token, error) {
	keyFunc := func(token *jwtpkg.Token) (interface{}, error) {
		_, ok := token.Method.(*jwtpkg.SigningMethodHMAC)
		if !ok {
			return nil, ErrTokenInvalid
		}
		return []byte(jwt.Key), nil
	}

	// parse jwt token
	jwtToken, err := jwtpkg.ParseWithClaims(token, &JWTCustomClaims{}, keyFunc)
	return jwtToken, err
}

// parse token for outside
func (jwt *JWT) VerifyToken(c *gin.Context) (*JWTCustomClaims, error) {
	// get tokenStr from Header
	tokenStr, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}

	// parse token
	token, jwtErr := jwt.parseToken(tokenStr)

	if jwtErr != nil {
		// token expired
		if errors.Is(jwtErr, jwtpkg.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}

		// malformed token
		if errors.Is(jwtErr, jwtpkg.ErrTokenMalformed) {
			return nil, ErrTokenMalformed
		}

		// invalid token
		return nil, ErrTokenInvalid
	}

	// if ok then return claims
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}

	// parse error then return error message
	return nil, ErrTokenInvalid
}

// refresh token
func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {
	// get tokenStr from Header
	tokenStr, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return "", parseErr
	}

	// parse token
	token, jwtErr := jwt.parseToken(tokenStr)
	if jwtErr != nil {
		// if error is not expired then return error
		if !errors.Is(jwtErr, jwtpkg.ErrTokenExpired) {
			return "", jwtErr
		}
	}

	// parse JWTCustomClaims
	claims := token.Claims.(*JWTCustomClaims)

	// check if IssuedAt is over max refresh time
	maxTime := jwt.getTimeWithLocation().Add(-jwt.MaxRefresh).Unix()
	if claims.IssuedAt.Unix() > maxTime {
		if errors.Is(jwtErr, jwtpkg.ErrTokenExpired) {
			claims.RegisteredClaims.ExpiresAt = jwtpkg.NewNumericDate(jwt.getTimeWithLocation().Add(time.Duration(cast.ToInt64(config.Config.JWT.EXPIRETIME)) * time.Minute))
			return jwt.createToken(*claims)
		} else {
			return "", nil
		}
	}

	return "", ErrTokenExpiredMaxRefresh
}
