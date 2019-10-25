package migrations

import "go.jinya.de/ontheroad/database"

type Migration20191025T234219766Z struct {
}

func (migration *Migration20191025T234219766Z) GetVersion() string {
	return "20191025T234219766Z"
}

func (migration *Migration20191025T234219766Z) Execute() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(database.ConfigurationTable)
	if err != nil {
		panic(err)
	}
}
