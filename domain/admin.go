package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/j23063519/clean_architecture/pkg/hash"
	"gorm.io/gorm"
)

type Admin struct {
	UuidModel

	Account  string `json:"account" gorm:"uniqueIndex;comment:account;"`
	Password string `json:"password" gorm:"uniqueIndex;comment:password;"`

	CommonTimestmp
}

func (admin *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	if admin.ID == "" {
		admin.ID = uuid.NewString()
	}

	if !hash.BcryptIsHashed(admin.Password) {
		admin.Password = hash.BcryptHash(admin.Password)
	}
	return
}

func (admin *Admin) BeforeUpdate(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(admin.Password) {
		admin.Password = hash.BcryptHash(admin.Password)
	}
	return
}

type AdminUC interface {
	Register(c *gin.Context, req LoginAndRegisterRequest) (Admin, error)
	Login(c *gin.Context, req LoginAndRegisterRequest) (Admin, error)
	GetAdminByID(c *gin.Context, id string) (Admin, error)
}

type AdminRepo interface {
	CreateAdmin(c *gin.Context, admin *Admin) error
	GetAdmin(c *gin.Context) (Admin, error)
	GetAdminByAccount(c *gin.Context, account string) (Admin, error)
	GetAdminByID(c *gin.Context, id string) (Admin, error)
}
