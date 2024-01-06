package repositery

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/j23063519/clean_architecture/domain"
	"github.com/j23063519/clean_architecture/pkg/log"
	"gorm.io/gorm"
)

type adminRepo struct {
	ORM *gorm.DB
	SQL *sql.DB
}

func NewAdminRepo(gorm *gorm.DB, sql *sql.DB) domain.AdminRepo {
	return &adminRepo{
		ORM: gorm,
		SQL: sql,
	}
}

func (a *adminRepo) CreateAdmin(c *gin.Context, admin *domain.Admin) (err error) {
	err = a.ORM.Create(&admin).Error
	return
}

func (a *adminRepo) GetAdmin(c *gin.Context) (admin domain.Admin, err error) {
	admin, ok := c.MustGet("current_admin").(domain.Admin)
	if !ok {
		err = errors.New("admin not exist")
		log.ErrorJSON("adminRepo", "GetAdmin", err)
		return admin, err
	}

	return admin, nil
}

func (a *adminRepo) GetAdminByAccount(c *gin.Context, account string) (admin domain.Admin, err error) {
	err = a.ORM.Where("account = ?", account).First(&admin).Error
	return
}

func (a *adminRepo) GetAdminByID(c *gin.Context, id string) (admin domain.Admin, err error) {
	err = a.ORM.Where("id = ?", id).First(&admin).Error
	return
}
