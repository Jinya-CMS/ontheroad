package migrations

import "go.jinya.de/ontheroad/database"

type Migration20191009T215919766Z struct {
}

func (migration *Migration20191009T215919766Z) GetVersion() string {
	return "20191009T215919766Z"
}

func (migration *Migration20191009T215919766Z) Execute() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE UNIQUE INDEX user_email_uindex ON \"user\" (email);")
	if err != nil {
		panic(err)
	}
}
