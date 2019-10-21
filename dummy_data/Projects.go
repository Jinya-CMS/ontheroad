package dummy_data

import "go.jinya.de/ontheroad/database"

var TestProject = database.Project{
	Id:              "bcdd3304-7a73-4a85-a2d6-f62d130ee932",
	Name:            "Jinya Gallery CMS",
	YouTrackServer:  "https://jinya.myjetbrains.com/youtrack/",
	VersionsQuery:   "api/admin/customFieldSettings/bundles/version/71-2",
	SubsystemsQuery: "api/admin/customFieldSettings/bundles/enum/66-9",
	TypesQuery:      "api/admin/customFieldSettings/bundles/enum/66-7",
	Key:             "JGCMS",
}

func ClearProjects() {
	connection, err := database.Connect()
	if err != nil {
		panic(err)
	}

	defer connection.Close()

	_, err = connection.Exec("DELETE FROM \"project\" WHERE id = $1", TestProject.Id)

	if err != nil {
		panic(err)
	}
}

func FillProjects() {
	connection, err := database.Connect()
	if err != nil {
		panic(err)
	}

	defer connection.Close()

	_, err = connection.Exec("INSERT INTO \"project\" (id, name, youtrackServer, key, versionsQuery, typesQuery, subsystemsQuery) VALUES ($1, $2, $3, $4, $5, $6, $7)", TestProject.Id, TestProject.Name, TestProject.YouTrackServer, TestProject.Key, TestProject.VersionsQuery, TestProject.TypesQuery, TestProject.SubsystemsQuery)

	if err != nil {
		panic(err)
	}
}
