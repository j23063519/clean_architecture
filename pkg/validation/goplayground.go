package validation

import (
	"github.com/j23063519/clean_architecture/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func customMessage(errs validator.ValidationErrors) string {
	if len(errs) > 0 {
		return errs[0].Error()
	}
	return "no errors"
}

func errHandle(c *gin.Context, err error) bool {
	if err != nil {
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Response(c, 400, err.Error(), nil)
			return false
		}

		response.Response(c, 422, customMessage(verrs), nil)
		return false
	}

	return true
}

func ValidateByGoPlayground(c *gin.Context, obj interface{}) bool {
	return errHandle(c, c.ShouldBind(obj))
}

func ValidateUriByGoPlayground(c *gin.Context, obj interface{}) bool {
	return errHandle(c, c.ShouldBindUri(obj))
}
