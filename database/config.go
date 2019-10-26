package database

import (
	"github.com/mholt/binding"
	"net/http"
)

type Configuration struct {
	Id    string
	Key   string
	Value string
}

func (config *Configuration) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&config.Value: binding.Field{
			Form:         "value",
			Required:     true,
			ErrorMessage: "The value is required",
		},
		&config.Key: binding.Field{
			Form:         "key",
			Required:     true,
			ErrorMessage: "The key is required",
		},
	}
}

// language=sql
var ConfigurationTable = `
CREATE TABLE Configuration (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    key TEXT NOT NULL UNIQUE,
    value TEXT NOT NULL
)`

func GetConfiguration(key string) (*Configuration, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	config := new(Configuration)
	err = db.Get(config, "SELECT key AS Key, value AS VALUE, id AS Id FROM Configuration WHERE Key = $1", key)

	return config, err
}

func GetConfigurations() ([]Configuration, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	config := new([]Configuration)
	err = db.Select(config, "SELECT key AS Key, value AS Value, id AS Id FROM Configuration")

	return *config, err
}

func SetConfiguration(key string, value string) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	count := new(int)
	err = db.Get(count, "SELECT COUNT(*) FROM configuration WHERE key = $1", key)
	if err != nil {
		return err
	}

	if *count == 0 {
		_, err = db.Exec("INSERT INTO Configuration (key, value) VALUES ($1, $2)", key, value)

		return err
	} else {
		_, err = db.Exec("UPDATE configuration SET value = $1 WHERE key = $2", value, key)

		return err
	}
}

func DeleteConfiguration(key string) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM Configuration WHERE key = $1", key)

	return nil
}
