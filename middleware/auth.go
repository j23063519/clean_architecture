package middleware

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	adminRepo "github.com/j23063519/clean_architecture/api/v1/admin/repositery"
	userRepo "github.com/j23063519/clean_architecture/api/v1/user/repositery"
	"github.com/j23063519/clean_architecture/pkg/database"
	"github.com/j23063519/clean_architecture/pkg/jwt"
	"github.com/j23063519/clean_architecture/pkg/response"

	"gorm.io/gorm"
)

func Auth(tpe string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 從 header Authorization:Bearer xxxxx 獲取訊息，並驗證 JWT 的準確性
		claims, err := jwt.NewJWT().VerifyToken(c)

		// 解析失敗
		if err != nil {
			response.Response(c, 401, err.Error(), nil)
			return
		}

		// 獲取 db
		currentDB := database.DBs["default"]
		// 獲取失敗
		if currentDB.Gorm == nil || currentDB.Sql == nil {
			response.Response(c, 500, "database can't connect", nil)
			return
		}

		switch tpe {
		case "user":
			if ok := user(c, claims, currentDB.Gorm, currentDB.Sql); !ok {
				return
			}
		case "admin":
			if ok := admin(c, claims, currentDB.Gorm, currentDB.Sql); !ok {
				return
			}
		default:
			response.Response(c, 500, "auth type not found", nil)
			return
		}

		c.Next()
	}
}

// 用戶者端
func user(c *gin.Context, claims *jwt.JWTCustomClaims, gorm *gorm.DB, sql *sql.DB) bool {
	// 獲取 用戶者 訊息
	userRepo := userRepo.NewUserRepo(gorm, sql)
	userModel, _ := userRepo.GetUserByID(c, claims.UserID)

	// 用戶者未找到
	if userModel.ID == "" {
		response.Response(c, 401, "user not found", nil)
		return false
	}

	// 儲存用戶者訊息
	c.Set("current_table", "user")
	c.Set("current_user_id", userModel.ID)
	c.Set("current_user_name", userModel.Account)
	c.Set("current_user", userModel)

	return true
}

// 管理者端
func admin(c *gin.Context, claims *jwt.JWTCustomClaims, gorm *gorm.DB, sql *sql.DB) bool {
	// 獲取 管理者 訊息
	repo := adminRepo.NewAdminRepo(gorm, sql)
	adminModel, _ := repo.GetAdminByID(c, claims.UserID)

	// 管理者未找到
	if adminModel.ID == "" || adminModel.Password == "" {
		response.Response(c, 401, "admin not found", nil)
		return false
	}

	// 儲存管理者訊息
	c.Set("current_table", "admin")
	c.Set("current_admin_id", adminModel.ID)
	c.Set("current_admin_name", adminModel.Account)
	c.Set("current_admin", adminModel)

	return true
}
