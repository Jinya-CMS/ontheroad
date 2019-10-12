package migrations

import "go.jinya.de/ontheroad/database"

type Migration20191012T003719766Z struct {
}

func (migration *Migration20191012T003719766Z) GetVersion() string {
	return "20191012T003719766Z"
}

func (migration *Migration20191012T003719766Z) Execute() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("ALTER TABLE project drop column query")
	if err != nil {
		panic(err)
	}
}
