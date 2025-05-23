package init

import (
	logpkg "github.com/j23063519/clean_architecture/pkg/log"

	"github.com/j23063519/clean_architecture/config"
	"github.com/j23063519/clean_architecture/pkg/console"
	"github.com/j23063519/clean_architecture/pkg/database"
	"gorm.io/gorm"
)

func SetDB() {
	configs := make(map[string]gorm.Dialector)

	// container
	configs[config.Config.DB.PGSQL.DATABASE] = database.PGSQLDialector(
		config.Config.DB.PGSQL.HOST,
		config.Config.DB.PGSQL.USERNAME,
		config.Config.DB.PGSQL.PASSWORD,
		config.Config.DB.PGSQL.DATABASE,
		"5432",
		config.Config.App.TIMEZONE,
	)

	if err := database.Connect(configs, logpkg.NewGormLogger()); err != nil {
		logpkg.ErrorJSON("Database", "db connection", err)
		console.Error("Couldn't connect to database: " + err.Error())
	}
}
