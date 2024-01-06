package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/j23063519/clean_architecture/pkg/hash"
	"gorm.io/gorm"
)

type User struct {
	UuidModel

	Account  string `json:"account" gorm:"uniqueIndex;comment:account;"`
	Password string `json:"password" gorm:"uniqueIndex;comment:password;"`

	CommonTimestmp
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.ID == "" {
		user.ID = uuid.NewString()
	}

	if !hash.BcryptIsHashed(user.Password) {
		user.Password = hash.BcryptHash(user.Password)
	}
	return
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(user.Password) {
		user.Password = hash.BcryptHash(user.Password)
	}
	return
}

type UserUC interface {
	Register(c *gin.Context, req LoginAndRegisterRequest) (User, error)
	Login(c *gin.Context, req LoginAndRegisterRequest) (User, error)
	GetUserByID(c *gin.Context, id string) (User, error)
}

type UserRepo interface {
	CreateUser(c *gin.Context, user *User) error
	GetUser(c *gin.Context) (User, error)
	GetUserByAccount(c *gin.Context, account string) (User, error)
	GetUserByID(c *gin.Context, id string) (User, error)
}
