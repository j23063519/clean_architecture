package repositery

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/j23063519/clean_architecture/domain"
	"github.com/j23063519/clean_architecture/pkg/log"
	"gorm.io/gorm"
)

type userRepo struct {
	ORM *gorm.DB
	SQL *sql.DB
}

func NewUserRepo(gorm *gorm.DB, sql *sql.DB) domain.UserRepo {
	return &userRepo{
		ORM: gorm,
		SQL: sql,
	}
}

func (u *userRepo) CreateUser(c *gin.Context, user *domain.User) (err error) {
	err = u.ORM.Create(&user).Error
	return
}

func (u *userRepo) GetUser(c *gin.Context) (user domain.User, err error) {
	user, ok := c.MustGet("current_user").(domain.User)
	if !ok {
		err = errors.New("user not exist")
		log.ErrorJSON("userRepo", "GetUser", err)
		return user, err
	}

	return user, nil
}

func (u *userRepo) GetUserByAccount(c *gin.Context, account string) (user domain.User, err error) {
	err = u.ORM.Where("account = ?", account).First(&user).Error
	return
}

func (u *userRepo) GetUserByID(c *gin.Context, id string) (user domain.User, err error) {
	err = u.ORM.Where("id = ?", id).First(&user).Error
	return
}
