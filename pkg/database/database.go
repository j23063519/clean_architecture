package database

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type db struct {
	Gorm *gorm.DB
	Sql  *sql.DB
}

var DBs map[string]db

// connect database
func Connect(configs map[string]gorm.Dialector, _logger gormlogger.Interface) error {
	DBs = make(map[string]db)

	for k, v := range configs {
		// gorm connect
		gorm, err := gorm.Open(v, &gorm.Config{
			Logger: _logger,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
		if err != nil {
			return err
		}

		// get base sql
		sql, err := gorm.DB()
		setConnectionPool(sql)

		if err != nil {
			return err
		}

		DBs[k] = db{gorm, sql}
	}

	return nil
}

// set connection pool
func setConnectionPool(sql *sql.DB) {
	// Used to set the maximum number of idle connections in the connection pool
	sql.SetMaxIdleConns(10)

	// Set the maximum number of open database connections
	sql.SetMaxOpenConns(100)

	// Sets the maximum time a connection can be reused
	sql.SetConnMaxLifetime(time.Hour)
}

// PGSQLDialector gives a postgresql gorm.Dialector
func PGSQLDialector(host, username, password, database, port, timezone string) gorm.Dialector {
	return postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v",
			host,
			username,
			password,
			database,
			port,
			timezone,
		),
	})
}

// MYSQLDialector gives a mysql/marizfb gorm.Dialector
func MYSQLDialector(host, username, password, database, port, charset, timezone string) gorm.Dialector {
	tz := "Asia%2fTaipei"
	if timezone != "" {
		tz = timezone
	}

	return mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=%v",
			username,
			password,
			host,
			port,
			database,
			charset,
			tz,
		),
	})
}
