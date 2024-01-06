package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/j23063519/clean_architecture/pkg/validation"
	"github.com/thedevsaddam/govalidator"
)

func AttachNecessary(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		// size unit:bytes
		// - 1024 bytes is 1kb
		// - 1048576 bytes is 1mb
		// - 5242880 bytes is 5mb
		// - 10485760 bytes is 10mb
		// - 20971520 bytes is 20mb
		// - 52428800 bytes is 50mb
		"file:attach": []string{"ext:png,jpg,jpeg", "size:5242880", "required"},
	}
	messages := govalidator.MapData{
		"file:attach": []string{
			"ext:The attach is invalid",
			"size:The size of the attach must be smaller than 5MB",
			"required:The attach is required",
		},
	}
	// start validate
	return validation.ValidateFile(c, data, rules, messages)
}

func AttachDefault(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		// size unit:bytes
		// - 1024 bytes is 1kb
		// - 1048576 bytes is 1mb
		// - 5242880 bytes is 5mb
		// - 10485760 bytes is 10mb
		// - 20971520 bytes is 20mb
		// - 52428800 bytes is 50mb
		"file:attach": []string{"ext:png,jpg,jpeg", "size:5242880"},
	}
	messages := govalidator.MapData{
		"file:attach": []string{
			"ext:The attach is invalid",
			"size:The size of the attach must be smaller than 5MB",
		},
	}
	// start validate
	return validation.ValidateFile(c, data, rules, messages)
}

type LoginAndRegisterRequest struct {
	Account  string `json:"account" binding:"required" example:"account"`
	Password string `json:"password" binding:"required" example:"password"`
}

type Type interface {
	User | Admin
}

type LoginResponse[T Type] struct {
	T     T
	Token string `json:"token"`
}
