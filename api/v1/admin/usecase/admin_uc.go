package usecase

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/j23063519/clean_architecture/domain"
	"github.com/j23063519/clean_architecture/pkg/hash"
	"github.com/j23063519/clean_architecture/pkg/response"
	"gorm.io/gorm"
)

type adminUC struct {
	adminRepo domain.AdminRepo
}

func NewAdminUC(adminRepo domain.AdminRepo) domain.AdminUC {
	return &adminUC{
		adminRepo: adminRepo,
	}
}

func (a *adminUC) Register(c *gin.Context, req domain.LoginAndRegisterRequest) (admin domain.Admin, err error) {
	admin = domain.Admin{
		Account:  req.Account,
		Password: req.Password,
	}

	err = a.adminRepo.CreateAdmin(c, &admin)
	if err != nil {
		response.Response(c, 500, err.Error(), nil)
		return
	}

	return
}

func (a *adminUC) Login(c *gin.Context, req domain.LoginAndRegisterRequest) (admin domain.Admin, err error) {
	admin, err = a.adminRepo.GetAdminByAccount(c, req.Account)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Admin{}, nil
		}
		response.Response(c, 500, err.Error(), nil)
		return
	}

	if !hash.BcryptCheck(req.Password, admin.Password) {
		err = errors.New("password not correct")
		response.Response(c, 500, err.Error(), nil)
		return
	}

	return
}

func (a *adminUC) GetAdminByID(c *gin.Context, id string) (admin domain.Admin, err error) {
	admin, err = a.adminRepo.GetAdminByID(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Admin{}, nil
		}
		response.Response(c, 500, err.Error(), nil)
		return
	}

	return
}
