package database

import (
	"database/sql"
	"fmt"
	"github.com/mholt/binding"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type User struct {
	Id            string
	Name          string
	Email         string
	Password      string
	TwoFactorCode sql.NullString
}

func (user *User) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&user.Name: binding.Field{
			Form:         "name",
			Required:     true,
			ErrorMessage: "The name is required",
		},
		&user.Email: binding.Field{
			Form:         "email",
			Required:     true,
			ErrorMessage: "The email is required",
		},
		&user.Password: binding.Field{
			Form:         "password",
			Required:     true,
			ErrorMessage: "The password is required",
		},
	}
}

// language=sql
var UserCreateTable = `
CREATE TABLE "user" (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    email text UNIQUE NOT NULL,
    password text NOT NULL,
    twoFactorCode text
)`

func GetUsers(offset int, limit int, keyword string) ([]User, int, error) {
	db, err := Connect()
	if err != nil {
		return nil, -1, err
	}

	defer db.Close()
	// language=sql prefix=SELECT * FROM user
	whereClause := "WHERE name ilike $1 OR email ilike $1"
	// language=sql
	selectQuery := fmt.Sprintf("SELECT u.Id AS Id, u.name AS Name, u.email AS Email, u.password AS Password, u.twoFactorCode AS TwoFactorCode FROM \"user\" u %s LIMIT $2 OFFSET $3", whereClause)

	users := new([]User)

	err = db.Select(users, selectQuery, "%"+keyword+"%", limit, offset)
	if err != nil {
		return nil, -1, err
	}

	// language=sql
	countQuery := fmt.Sprintf("SELECT COUNT(u.id) FROM \"user\" u %s", whereClause)
	totalCount := new(int)
	err = db.Get(totalCount, countQuery, "%"+keyword+"%")

	return *users, *totalCount, err
}

func GetUser(id string) (*User, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	user := new(User)
	err = db.Get(user, "SELECT * FROM \"user\" WHERE id = $1", id)

	return user, err
}

func CreateUser(user *User) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	defer db.Close()
	password, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO \"user\" (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, password)

	return err
}

func UpdateUser(user *User) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	defer db.Close()
	_, err = db.Exec("UPDATE \"user\" SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, user.Id)

	return err
}

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 13)

	return string(hashed), err
}

func ValidateEmailAndPassword(email string, password string) (*User, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	user := new(User)

	err = db.Get(user, "SELECT * FROM \"user\" WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}
