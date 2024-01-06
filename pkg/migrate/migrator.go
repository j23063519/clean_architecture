package migrate

import (
	"os"

	"github.com/j23063519/clean_architecture/pkg/console"
	"github.com/j23063519/clean_architecture/pkg/database"
	"github.com/j23063519/clean_architecture/pkg/util"
	"gorm.io/gorm"
)

type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
	DBName   string
}

// Migration corresponds to a piece of data in the migrations table of the data.
type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

// NewMigrator creates a Migrator instance to perform migration operations
func NewMigrator(dbName string) *Migrator {
	// must has db
	if database.DBs[dbName].Gorm == nil || database.DBs[dbName].Sql == nil {
		console.Error("Can't find DB: " + dbName)
		return nil
	}

	// init
	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DBs[dbName].Gorm,
		Migrator: database.DBs[dbName].Gorm.Migrator(),
		DBName:   dbName,
	}

	// create a migration table
	migrator.createMigrationsTable()

	return migrator
}

// create migrations table
func (migrator *Migrator) createMigrationsTable() {
	// if migrations table not exist then create it
	if !migrator.Migrator.HasTable(&Migration{}) {
		migrator.Migrator.CreateTable(&Migration{})
	}
}

// Execute all unmigrated files
func (migrator *Migrator) Up() {
	// Read all migration files, making sure they are sorted by time
	migrateFiles := migrator.readAllMigrationFiles()

	// get current batch value
	batch := migrator.getBatch()

	// get all migration datas
	migrations := []Migration{}
	migrator.DB.Find(&migrations)

	// Determine whether the database is up to date based on runed
	runed := false

	// Traverse the migration file. If it has not been executed before, execute up
	for _, mfile := range migrateFiles {
		// Compare the file names to see if they have been run before. If not, perform the migration.
		if mfile.isNotMigrated(migrations) {
			migrator.runUpMigration(mfile, batch)
			runed = true
		}
	}

	if !runed {
		console.Success("Database is up to date")
	}
}

// Read migration files from file directory, ensuring correct time ordering
func (migrator *Migrator) readAllMigrationFiles() (migrateFiles []MigrationFile) {
	// read databas/migrations/ all files in the directory
	// basically, they will be sorted by file name.
	files, err := os.ReadDir(migrator.Folder)
	if err != nil {
		console.Error("Read file is error: " + err.Error())
	}

	for _, v := range files {
		filename := util.FileNameWithoutExtension(v.Name())

		mfile := getMigrationFile(filename)

		if len(mfile.FileName) > 0 {
			migrateFiles = append(migrateFiles, mfile)
		}
	}

	return migrateFiles
}

// get current batch value
func (migrator *Migrator) getBatch() (batch int) {
	// initial value is 1
	batch = 1

	// get last migration
	lastMigration := Migration{}
	if err := migrator.DB.Order("id DESC").First(&lastMigration).Error; err != nil {
		console.Error(err.Error())
	}

	// if lastMigration.ID greater than 0 then add 1
	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}

	return
}

func (migrator *Migrator) runUpMigration(mfile MigrationFile, batch int) {
	// excute up's sql
	if mfile.Up != nil {
		console.Warn("nmigrating: " + mfile.FileName)
		mfile.Up(database.DBs[migrator.DBName].Gorm.Migrator(), database.DBs[migrator.DBName].Sql)
		console.Success("nmigrated: " + mfile.FileName)
	}

	// create migration
	err := migrator.DB.Create(&Migration{Migration: mfile.FileName, Batch: batch}).Error
	if err != nil {
		console.ExitIf(err)
	}
}

// roll back to the migration performed in the previous step
func (migrator *Migrator) Rollback() {
	// get the last migration data
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)
	migratnios := []Migration{}
	migrator.DB.Where("batch = ?", lastMigration.Batch).Order("id DESC").Find(&migratnios)

	// rollback the last batch of migrations
	if !migrator.rollbackMigrations(migratnios) {
		console.Success("[migratnios] table is empty, nothing to rollback")
	}
}

// roll back to the previous migration and execute the down method of migrations in order from largest to smallest.
func (migrator *Migrator) rollbackMigrations(migrations []Migration) bool {
	// whether the initialization actually performs the action of rolling back to the previous migration
	runed := false

	for _, v := range migrations {
		console.Warn("rollback start: " + v.Migration)

		mfile := getMigrationFile(v.Migration)

		if mfile.Down == nil {
			console.Error("rollback error: " + mfile.FileName + " down not found")
			return runed
		}

		// excute down
		mfile.Down(database.DBs[migrator.DBName].Gorm.Migrator(), database.DBs[migrator.DBName].Sql)

		// if success then delete the record
		migrator.DB.Delete(&v)

		runed = true
		console.Success("rollback finish: " + v.Migration)
	}

	return runed
}

// roll back all migration
func (migrator *Migrator) Reset() {
	migrations := []Migration{}

	// read files in order from largest to smallest
	migrator.DB.Order("id DESC").Find(&migrations)

	// roll back all migration
	if !migrator.rollbackMigrations(migrations) {
		console.Success("[migratnios] table is empty, nothing to rollback")
	}
}

// roll back all migration，then excute up again
func (migrator *Migrator) Refresh() {
	// roll back all migration
	migrator.Reset()

	// excute up
	migrator.Up()
}

// delete all tables then excute up
func (migrator *Migrator) Fresh() {
	dbname := migrator.currentDB()

	err := migrator.deleteAllTables()
	console.ExitIf(err)
	console.Success("clearup database: " + dbname)

	migrator.createMigrationsTable()
	console.Success("[migrations] table created.")

	migrator.Up()
}

// get current database
func (migrator *Migrator) currentDB() string {
	return migrator.DB.Migrator().CurrentDatabase()
}

// delete all tables
func (migrator *Migrator) deleteAllTables() error {
	dbname := migrator.currentDB()
	tables := []string{}

	// get all tables
	err := migrator.DB.Table("information_schema.tables").
		Where("table_schema = ?", dbname).
		Pluck("table_name", &tables).
		Error
	if err != nil {
		return err
	}

	// close foreign keys checking
	// migrator.DB.Exec("SET foreign_key_checks = 0;")               // mysql
	migrator.DB.Exec("SET session_replication_role = 'replica';") // pgsql

	// delete all tables
	for _, table := range tables {
		if err := migrator.DB.Migrator().DropTable(table); err != nil {
			return err
		}
	}

	// open foreign keys checking
	// migrator.DB.Exec("SET foreign_key_checks = 1;")              // mysql
	migrator.DB.Exec("SET session_replication_role = 'origin';") // pgsql

	return nil
}
