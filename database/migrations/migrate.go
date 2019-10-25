package migrations

import "go.jinya.de/ontheroad/database"

type Migration interface {
	Execute()
	GetVersion() string
}

// language=sql
var MigrationsTable = `
CREATE TABLE IF NOT EXISTS migrations (
    Version varchar(255) PRIMARY KEY
)
`

var Migrations = []Migration{
	new(Migration20191006T165019766Z),
	new(Migration20191008T213119766Z),
	new(Migration20191009T215919766Z),
	new(Migration20191010T200519766Z),
	new(Migration20191012T003719766Z),
	new(Migration20191012T203219766Z),
	new(Migration20191025T234219766Z),
}

func createMigrationsTable() error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec(MigrationsTable)
	if err != nil {
		return err
	}

	return nil
}

func saveMigration(version string) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()
	// language=sql
	_, err = db.Exec("INSERT INTO Migrations (version) VALUES ($1)", version)

	return err
}

func isMigrated(version string) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		return true, err
	}

	defer db.Close()

	var count int
	// language=sql
	err = db.Get(&count, "SELECT COUNT(*) FROM migrations WHERE version = $1", version)
	if err != nil {
		return true, err
	}

	if count == 1 {
		return true, nil
	}

	return false, err
}

func Migrate() error {
	err := createMigrationsTable()
	if err != nil {
		return err
	}

	for _, migration := range Migrations {
		version := migration.GetVersion()
		migrated, err := isMigrated(version)
		if err != nil {
			return err
		}

		if !migrated {
			migration.Execute()
			err = saveMigration(version)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
