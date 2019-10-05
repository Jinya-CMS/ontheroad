package database

func CreateDatabase() error {
	db, err := Connect()
	if err != nil {
		return err
	}

	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	if err != nil {
		return err
	}

	_, err = tx.Exec(UserCreateTable)
	if err != nil {
		return err
	}

	return tx.Commit()
}
