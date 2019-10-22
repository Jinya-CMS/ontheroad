package migrations

import "go.jinya.de/ontheroad/database"

type Migration20191012T203219766Z struct {
}

func (migration *Migration20191012T203219766Z) GetVersion() string {
	return "20191012T203219766Z"
}

func (migration *Migration20191012T203219766Z) Execute() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE UNIQUE INDEX project_key_uindex ON project (key)")
	if err != nil {
		panic(err)
	}
}
