package response

import "github.com/gin-gonic/gin"

// response structure
type RespnseStr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// response
func Response(c *gin.Context, code int, msg string, data interface{}) {
	switch code {
	case 200:
		success(c, code, data)
	case 201:
		success2(c, code, msg)
	case 400:
		badRequest(c, code, msg)
	case 401:
		unauthorized(c, code, msg)
	case 403:
		forbidden(c, code, msg)
	case 404:
		notFound(c, code, msg)
	case 422:
		unprocessableEntity(c, code, msg)
	case 429:
		tooManyRequest(c, code, msg)
	case 500:
		internalServerError(c, code, msg)
	default:
		internalServerError(c, code, msg)
	}
}

// json
func json(c *gin.Context, code int, msg string, data interface{}) {
	jsonData := RespnseStr{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	if code > 299 {
		c.AbortWithStatusJSON(code, jsonData)
	} else {
		c.JSON(code, jsonData)
	}
}

// if other => default: other，or use default message
func msgf(defualt, other string) string {
	if other != "" {
		return defualt + ": " + other
	}
	return defualt
}

// 200 default return success message and return data
func success(c *gin.Context, code int, data interface{}) {
	json(c, code, msgf("success", ""), data)
}

// 201 create / update / delete and return information
func success2(c *gin.Context, code int, msg string) {
	json(c, code, msgf(msg+"success", ""), nil)
}

// 400 Bad Request
func badRequest(c *gin.Context, code int, msg string) {
	json(c, code, msgf("bad request", msg), nil)
}

// 401 Unauthorized
func unauthorized(c *gin.Context, code int, msg string) {
	json(c, code, msgf("unauthorized", msg), nil)
}

// 403 Forbidden
func forbidden(c *gin.Context, code int, msg string) {
	json(c, code, msgf("forbidden", msg), nil)
}

// 404 not found
func notFound(c *gin.Context, code int, msg string) {
	json(c, code, msgf("404 not found", msg), nil)
}

// 422 Unprocessable Entity
func unprocessableEntity(c *gin.Context, code int, msg string) {
	json(c, code, msgf("unprocessable entity", msg), nil)
}

// 429 Too Many Requests
func tooManyRequest(c *gin.Context, code int, msg string) {
	json(c, code, msgf("too many requests", msg), nil)
}

// 500 Internal Server Error
func internalServerError(c *gin.Context, code int, msg string) {
	json(c, code, msgf("internal server error", msg), nil)
}
