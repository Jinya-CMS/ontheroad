package migrations

import "go.jinya.de/ontheroad/database"

type Migration20191006T165019766Z struct {
}

func (migation *Migration20191006T165019766Z) GetVersion() string {
	return "20191006T165019766Z"
}

func (migration *Migration20191006T165019766Z) Execute() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(database.AuthTokenTable)
	if err != nil {
		panic(err)
	}
}
