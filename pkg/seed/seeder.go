package seed

import (
	"github.com/j23063519/clean_architecture/pkg/console"
	"github.com/j23063519/clean_architecture/pkg/database"
	"gorm.io/gorm"
)

// saving all seeders
var seeders []Seeder

// Seeder array in order of execution
var orderedSeederNames []string

// excute gorm
type SeederFunc func(*gorm.DB)

// seeder corresponds to each Seeder file in the database/seeders directory
type Seeder struct {
	Func SeederFunc
	Name string
}

// add seeder to seeders
func Add(name string, fn SeederFunc) {
	seeders = append(seeders, Seeder{
		Name: name,
		Func: fn,
	})
}

// set up an array of Seeders in order of execution
func SetRunOrder(names []string) {
	orderedSeederNames = names
}

// get seeder by name
func GetSeeder(name string) Seeder {
	for _, v := range seeders {
		if name == v.Name {
			return v
		}
	}
	return Seeder{}
}

// excute all seeders
func RunAll(dbName string) {
	if len(dbName) < 1 || database.DBs[dbName].Gorm == nil {
		console.Error("Running All Seeders Fail: dbName not exist")
		return
	}

	// first excute ordered seeder
	executed := make(map[string]string)
	for _, name := range orderedSeederNames {
		sdr := GetSeeder(name)
		if len(sdr.Name) > 0 {
			console.Warn("Running Ordered Seeder: " + sdr.Name)
			sdr.Func(database.DBs[dbName].Gorm)
			executed[name] = name
		}
	}

	// excute other none ordered seeders
	for _, v := range seeders {
		if _, ok := executed[v.Name]; !ok {
			console.Warn("Running Seeder: " + v.Name)
			v.Func(database.DBs[dbName].Gorm)
		}
	}
}

// excute single seeder
func RunSeeder(name, dbName string) {
	if len(dbName) < 1 || database.DBs[dbName].Gorm == nil {
		console.Error("Running Seeder Fail: dbName not exist")
		return
	}

	for _, v := range seeders {
		if name == v.Name {
			console.Warn("Running Seeder: " + v.Name)
			v.Func(database.DBs[dbName].Gorm)
			break
		}
	}
}
