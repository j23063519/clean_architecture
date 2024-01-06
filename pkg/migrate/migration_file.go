package migrate

import (
	"database/sql"

	"gorm.io/gorm"
)

// migrationFunc define up and down
type migrationFunc func(gorm.Migrator, *sql.DB)

// migrationFiles
var migrationFiles []MigrationFile

// MigrationFile
type MigrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	FileName string
}

// Add a migration file
func Add(name string, up, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		Up:       up,
		Down:     down,
		FileName: name,
	})
}

// Get the Migration File by its name
func getMigrationFile(name string) MigrationFile {
	for _, v := range migrationFiles {
		if name == v.FileName {
			return v
		}
	}

	return MigrationFile{}
}

// isNotMigrated
func (mfile MigrationFile) isNotMigrated(migrations []Migration) bool {
	for _, migration := range migrations {
		if migration.Migration == mfile.FileName {
			return false
		}
	}
	return true
}
