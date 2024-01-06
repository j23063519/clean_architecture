package cmd

import (
	"fmt"

	"github.com/j23063519/clean_architecture/database/seeders"

	"github.com/j23063519/clean_architecture/pkg/console"
	"github.com/j23063519/clean_architecture/pkg/database"
	"github.com/j23063519/clean_architecture/pkg/seed"
	"github.com/spf13/cobra"
)

var CmdDBSeed = &cobra.Command{
	Use:   "seed",
	Short: "Insert Data To The Database",
	Run:   runSeeders,
	Args:  cobra.MaximumNArgs(1),
}

func init() {
	CmdDBSeed.Flags().StringVarP(&DBName, "dbName", "d", "", "Database name to Seed")
	CmdDBSeed.MarkFlagRequired("dbName")
}

func runSeeders(cmd *cobra.Command, args []string) {
	if database.DBs[DBName].Gorm == nil || database.DBs[DBName].Sql == nil {
		console.Error(fmt.Sprintf("Database [%v] not exist", DBName))
		return
	}

	seeders.Initialize()

	if len(args) > 0 {
		seederName := args[0]
		seeder := seed.GetSeeder(seederName)
		if len(seeder.Name) < 1 {
			console.Error("Seeder not found: " + seederName)
			return
		}

		seed.RunSeeder(seederName, DBName)
	} else {
		seed.RunAll(DBName)

		console.Success("Seeding Success")
	}
}
