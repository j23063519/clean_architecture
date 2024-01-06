package route

import (
	"github.com/j23063519/clean_architecture/docs"
	"github.com/j23063519/clean_architecture/pkg/database"

	adminHandler "github.com/j23063519/clean_architecture/api/v1/admin/delivery"
	adminRepo "github.com/j23063519/clean_architecture/api/v1/admin/repositery"
	adminUC "github.com/j23063519/clean_architecture/api/v1/admin/usecase"
	userHandler "github.com/j23063519/clean_architecture/api/v1/user/delivery"
	userRepo "github.com/j23063519/clean_architecture/api/v1/user/repositery"
	userUC "github.com/j23063519/clean_architecture/api/v1/user/usecase"

	"github.com/gin-gonic/gin"
	"github.com/j23063519/clean_architecture/config"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterHandler(router *gin.Engine) {
	// 共用 gorm sql
	gorm := database.DBs[config.Config.DB.PGSQL.DATABASE].Gorm
	sql := database.DBs[config.Config.DB.PGSQL.DATABASE].Sql

	// admin
	adminHandler.NewAdminHandler(
		router,
		adminUC.NewAdminUC(
			adminRepo.NewAdminRepo(gorm, sql),
		),
	)

	// user
	userHandler.NewUserHandler(
		router,
		userUC.NewUserUC(
			userRepo.NewUserRepo(gorm, sql),
		),
	)

	docs.SwaggerInfo.Title = "Swagger API"
	docs.SwaggerInfo.Description = "This is a api document of xxx."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.Config.App.HOST + ":" + config.Config.App.PORT
	docs.SwaggerInfo.BasePath = config.Config.App.BASEPATH
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.URL("")))
}
