package database

import (
	"database/sql"
	"fmt"
	"github.com/mholt/binding"
	"net/http"
)

type Project struct {
	Id              string
	Name            string
	Query           string
	YouTrackServer  string
	VersionsQuery   sql.NullString
	SubsystemsQuery sql.NullString
	TypesQuery      sql.NullString
	Key             string
}

func (project *Project) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&project.Name: binding.Field{
			Form:         "name",
			Required:     true,
			ErrorMessage: "The name is required",
		},
		&project.Query: binding.Field{
			Form:         "query",
			Required:     true,
			ErrorMessage: "The query is required",
		},
		&project.YouTrackServer: binding.Field{
			Form:         "youtrack_server",
			Required:     true,
			ErrorMessage: "The YouTrack server is required",
		},
		&project.SubsystemsQuery: binding.Field{
			Form:     "subsystems_query",
			Required: true,
		},
		&project.TypesQuery: binding.Field{
			Form:     "types_query",
			Required: true,
		},
		&project.VersionsQuery: binding.Field{
			Form:     "versions_query",
			Required: true,
		},
		&project.Key: binding.Field{
			Form:         "key",
			Required:     true,
			ErrorMessage: "The key is required",
		},
	}
}

// language=sql
var ProjectCreateTable = `
CREATE TABLE "project" (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    query text NOT NULL,
    youtrackServer text NOT NULL,
    versionsQuery text,
    subsystemsQuery text,
    typesQuery text,
    key text NOT NULL
)`

func GetProjects(offset int, limit int, keyword string) ([]Project, int, error) {
	db, err := Connect()
	if err != nil {
		return nil, -1, err
	}

	defer db.Close()
	// language=sql prefix="SELECT * FROM project "
	whereClause := "WHERE name ilike $1"
	// language=sql
	selectQuery := fmt.Sprintf("SELECT p.Id AS Id, p.name AS Name, p.key AS Key, p.query AS Query, p.youtrackserver AS YouTrackServer, p.versionsquery AS VersionsQuery, p.subsystemsQuery AS SubsystemsQuery, p.typesQuery AS TypesQuery, p.key AS Key FROM project p %s LIMIT $2 OFFSET $3", whereClause)

	projects := new([]Project)

	err = db.Select(projects, selectQuery, "%"+keyword+"%", limit, offset)
	if err != nil {
		return nil, -1, err
	}

	// language=sql
	countQuery := fmt.Sprintf("SELECT COUNT(u.id) FROM \"project\" u %s", whereClause)
	totalCount := new(int)
	err = db.Get(totalCount, countQuery, "%"+keyword+"%")

	return *projects, *totalCount, err
}

func GetProject(id string) (*Project, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	project := new(Project)
	err = db.Get(project, "SELECT * FROM \"project\" WHERE id = $1", id)

	return project, err
}

func CreateProject(project *Project) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("INSERT INTO \"project\" (name, youtrackServer, query, key) VALUES ($1, $2, $3, $4)", project.Name, project.YouTrackServer, project.Query, project.Key)

	return err
}

func UpdateProject(project *Project) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	defer db.Close()
	_, err = db.Exec("UPDATE \"project\" SET name = $1, youtrackServer = $2, query = $3, key = $4 WHERE id = $5", project.Name, project.YouTrackServer, project.Query, project.Key, project.Id)

	return err
}

func DeleteProject(id string) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("DELETE FROM \"project\" WHERE id = $1", id)

	return err
}
