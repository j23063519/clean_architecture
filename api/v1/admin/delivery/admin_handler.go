package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/j23063519/clean_architecture/domain"
	"github.com/j23063519/clean_architecture/middleware"
	"github.com/j23063519/clean_architecture/pkg/jwt"
	"github.com/j23063519/clean_architecture/pkg/response"
	"github.com/j23063519/clean_architecture/pkg/validation"
)

type AdminHandler struct {
	adminUC domain.AdminUC
}

func NewAdminHandler(r *gin.Engine, adminUC domain.AdminUC) {
	handler := &AdminHandler{
		adminUC: adminUC,
	}

	v1 := r.Group("/api/v1")
	{
		// admin
		adminGroup := v1.Group("/admin")
		{
			adminGroup.POST("/register", handler.Register)
			adminGroup.POST("/login", handler.Login)
			adminGroup.GET("/info/:id", middleware.Auth("admin"), handler.Info)
		}
	}
}

// Register godoc
//
//	@Summary		register admin
//	@Description	register admin
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			request	body		domain.LoginAndRegisterRequest			true	"LoginAndRegisterRequest"
//	@success		200		{object}	response.RespnseStr{data=domain.Admin}	"success"
//	@failure		401		{object}	response.RespnseStr{data=object}		"unauthorized"
//	@failure		500		{object}	response.RespnseStr{data=object}		"system error"
//	@Router			/admin/register [post]
func (a *AdminHandler) Register(c *gin.Context) {
	req := domain.LoginAndRegisterRequest{}
	if !validation.ValidateByGoPlayground(c, &req) {
		return
	}

	admin, err := a.adminUC.Register(c, req)
	if err != nil {
		return
	}

	response.Response(c, 200, "", admin)
}

// Login godoc
//
//	@Summary		login admin
//	@Description	login admin
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			request	body		domain.LoginAndRegisterRequest									true	"LoginAndRegisterRequest"
//	@success		200		{object}	response.RespnseStr{data=domain.LoginResponse[domain.Admin]}	"success"
//	@failure		401		{object}	response.RespnseStr{data=object}								"unauthorized"
//	@failure		500		{object}	response.RespnseStr{data=object}								"system error"
//	@Router			/admin/login [post]
func (a *AdminHandler) Login(c *gin.Context) {
	req := domain.LoginAndRegisterRequest{}
	if !validation.ValidateByGoPlayground(c, &req) {
		return
	}

	admin, err := a.adminUC.Login(c, req)
	if err != nil {
		return
	}

	token := jwt.NewJWT().IssueToken("admin", admin.ID, admin.Account)

	response.Response(c, 200, "", domain.LoginResponse[domain.Admin]{
		T:     admin,
		Token: token,
	})
}

// Info godoc
//
//	@Summary		login admin
//	@Description	login admin
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string									true	"id"
//	@success		200	{object}	response.RespnseStr{data=domain.Admin}	"success"
//	@failure		401	{object}	response.RespnseStr{data=object}		"unauthorized"
//	@failure		500	{object}	response.RespnseStr{data=object}		"system error"
//	@Router			/admin/info/{id} [get]
func (a *AdminHandler) Info(c *gin.Context) {
	admin, err := a.adminUC.GetAdminByID(c, c.Param("id"))
	if err != nil {
		return
	}

	response.Response(c, 200, "", admin)
}
