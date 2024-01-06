package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/j23063519/clean_architecture/domain"
	"github.com/j23063519/clean_architecture/middleware"
	"github.com/j23063519/clean_architecture/pkg/jwt"
	"github.com/j23063519/clean_architecture/pkg/response"
	"github.com/j23063519/clean_architecture/pkg/validation"
)

type UserHandler struct {
	userUC domain.UserUC
}

func NewUserHandler(r *gin.Engine, userUC domain.UserUC) {
	handler := &UserHandler{
		userUC: userUC,
	}

	v1 := r.Group("/api/v1")
	{
		// user
		userGroup := v1.Group("/user")
		{
			userGroup.POST("/register", handler.Register)
			userGroup.POST("/login", handler.Login)
			userGroup.GET("/info/:id", middleware.Auth("user"), handler.Info)
		}
	}
}

// Register godoc
//
//	@Summary		register user
//	@Description	register user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body		domain.LoginAndRegisterRequest			true	"LoginAndRegisterRequest"
//	@success		200		{object}	response.RespnseStr{data=domain.User}	"success"
//	@failure		401		{object}	response.RespnseStr{data=object}		"unauthorized"
//	@failure		500		{object}	response.RespnseStr{data=object}		"system error"
//	@Router			/user/register [post]
func (u *UserHandler) Register(c *gin.Context) {
	req := domain.LoginAndRegisterRequest{}
	if !validation.ValidateByGoPlayground(c, &req) {
		return
	}

	user, err := u.userUC.Register(c, req)
	if err != nil {
		return
	}

	response.Response(c, 200, "", user)
}

// Login godoc
//
//	@Summary		login user
//	@Description	login user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body		domain.LoginAndRegisterRequest								true	"LoginAndRegisterRequest"
//	@success		200		{object}	response.RespnseStr{data=domain.LoginResponse[domain.User]}	"success"
//	@failure		401		{object}	response.RespnseStr{data=object}							"unauthorized"
//	@failure		500		{object}	response.RespnseStr{data=object}							"system error"
//	@Router			/user/login [post]
func (u *UserHandler) Login(c *gin.Context) {
	req := domain.LoginAndRegisterRequest{}
	if !validation.ValidateByGoPlayground(c, &req) {
		return
	}

	user, err := u.userUC.Login(c, req)
	if err != nil {
		return
	}

	token := jwt.NewJWT().IssueToken("user", user.ID, user.Account)

	response.Response(c, 200, "", domain.LoginResponse[domain.User]{
		T:     user,
		Token: token,
	})
}

// Info godoc
//
//	@Summary		login user
//	@Description	login user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string									true	"id"
//	@success		200	{object}	response.RespnseStr{data=domain.User}	"success"
//	@failure		401	{object}	response.RespnseStr{data=object}		"unauthorized"
//	@failure		500	{object}	response.RespnseStr{data=object}		"system error"
//	@Router			/user/info/{id} [get]
func (u *UserHandler) Info(c *gin.Context) {
	user, err := u.userUC.GetUserByID(c, c.Param("id"))
	if err != nil {
		return
	}

	response.Response(c, 200, "", user)
}
