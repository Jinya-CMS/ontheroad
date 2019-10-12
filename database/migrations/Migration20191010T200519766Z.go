package migrations

import "go.jinya.de/ontheroad/database"

type Migration20191010T200519766Z struct {
}

func (migration *Migration20191010T200519766Z) GetVersion() string {
	return "20191010T200519766Z"
}

func (migration *Migration20191010T200519766Z) Execute() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(database.ProjectCreateTable)
	if err != nil {
		panic(err)
	}
}
