package migrations

import "go.jinya.de/ontheroad/database"

type Migration20191008T213119766Z struct {
}

func (migation *Migration20191008T213119766Z) GetVersion() string {
	return "20191008T213119766Z"
}

func (migration *Migration20191008T213119766Z) Execute() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("ALTER TABLE \"auth_token\" drop column ip_address")
	if err != nil {
		panic(err)
	}
}
