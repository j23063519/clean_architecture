package cmd

import (
	"github.com/j23063519/clean_architecture/database/migrations"
	"github.com/j23063519/clean_architecture/pkg/migrate"
	"github.com/spf13/cobra"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
}

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run unmigrated migrations",
	Run:   runUp,
}

var CmdMigrateRollback = &cobra.Command{
	Use: "down",
	// set alias: migrae down == migrate rollback
	Aliases: []string{"rollback"},
	Short:   "Reverse the up command",
	Run:     runDown,
}

var CmdMigrateReset = &cobra.Command{
	Use:   "reset",
	Short: "Rollback all database migrations",
	Run:   runReset,
}

var CmdMigrateRefresh = &cobra.Command{
	Use:   "refresh",
	Short: "Reset and re-run all migrations",
	Run:   runRefresh,
}

var CmdMigrateFresh = &cobra.Command{
	Use:   "fresh",
	Short: "Drop all tables and re-run all migrations",
	Run:   runFresh,
}

// DBName only use up/rollback/reset/refresh/fresh --dbName
var DBName string

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
		CmdMigrateRollback,
		CmdMigrateReset,
		CmdMigrateRefresh,
		CmdMigrateFresh,
	)

	// up
	CmdMigrateUp.Flags().StringVarP(&DBName, "dbName", "d", "", "Database name to migrate up")
	CmdMigrateUp.MarkFlagRequired("dbName")
	// rollback
	CmdMigrateRollback.Flags().StringVarP(&DBName, "dbName", "d", "", "Database name to migrate up")
	CmdMigrateRollback.MarkFlagRequired("dbName")
	// reset
	CmdMigrateReset.Flags().StringVarP(&DBName, "dbName", "d", "", "Database name to migrate up")
	CmdMigrateReset.MarkFlagRequired("dbName")
	// refresh
	CmdMigrateRefresh.Flags().StringVarP(&DBName, "dbName", "d", "", "Database name to migrate up")
	CmdMigrateRefresh.MarkFlagRequired("dbName")
	// fresh
	CmdMigrateFresh.Flags().StringVarP(&DBName, "dbName", "d", "", "Database name to migrate up")
	CmdMigrateFresh.MarkFlagRequired("dbName")
}

func migrator() *migrate.Migrator {
	// execute database/migrations all migration files
	migrations.Initialize()

	// new migrator
	return migrate.NewMigrator(DBName)
}

func runUp(cmd *cobra.Command, args []string) {
	if migrator() != nil {
		migrator().Up()
	}
}

func runDown(cmd *cobra.Command, args []string) {
	if migrator() != nil {
		migrator().Rollback()
	}
}

func runReset(cmd *cobra.Command, args []string) {
	if migrator() != nil {
		migrator().Reset()
	}
}

func runRefresh(cmd *cobra.Command, args []string) {
	if migrator() != nil {
		migrator().Refresh()
	}
}

func runFresh(cmd *cobra.Command, args []string) {
	if migrator() != nil {
		migrator().Fresh()
	}
}
