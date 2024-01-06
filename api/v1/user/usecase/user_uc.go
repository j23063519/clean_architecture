package usecase

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/j23063519/clean_architecture/domain"
	"github.com/j23063519/clean_architecture/pkg/hash"
	"github.com/j23063519/clean_architecture/pkg/response"
	"gorm.io/gorm"
)

type userUC struct {
	userRepo domain.UserRepo
}

func NewUserUC(userRepo domain.UserRepo) domain.UserUC {
	return &userUC{
		userRepo: userRepo,
	}
}

func (u *userUC) Register(c *gin.Context, req domain.LoginAndRegisterRequest) (user domain.User, err error) {
	user = domain.User{
		Account:  req.Account,
		Password: req.Password,
	}

	err = u.userRepo.CreateUser(c, &user)
	if err != nil {
		response.Response(c, 500, err.Error(), nil)
		return
	}

	return
}

func (u *userUC) Login(c *gin.Context, req domain.LoginAndRegisterRequest) (user domain.User, err error) {
	user, err = u.userRepo.GetUserByAccount(c, req.Account)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, nil
		}
		response.Response(c, 500, err.Error(), nil)
		return
	}

	if !hash.BcryptCheck(req.Password, user.Password) {
		err = errors.New("password not correct")
		response.Response(c, 500, err.Error(), nil)
		return
	}

	return
}

func (u *userUC) GetUserByID(c *gin.Context, id string) (user domain.User, err error) {
	user, err = u.userRepo.GetUserByID(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, nil
		}
		response.Response(c, 500, err.Error(), nil)
		return
	}

	return
}
