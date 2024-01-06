package validation

import (
	"sort"

	"github.com/j23063519/clean_architecture/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

var errVal string

// ValidatorFunc
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func ValidateByGoValidator(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	// parse request, supported JSON、FORM and URL Query
	if err := c.ShouldBind(obj); err != nil {
		response.Response(c, 400, err.Error(), nil)
		return false
	}

	// form validation
	errs := handler(obj, c)

	// sort errs
	keys := make([]string, 0, len(errs))
	for k := range errs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		errVal = errs[key][0]
	}

	// only show one error
	if len(errs) > 0 {
		response.Response(c, 422, errVal, nil)
		return false
	}

	return true
}

func ValidateForRequest(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {

	// init setting
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}

	// start validate
	return govalidator.New(opts).ValidateStruct()
}

func ValidateFile(c *gin.Context, data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Request:       c.Request,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	// start validate
	return govalidator.New(opts).Validate()
}
