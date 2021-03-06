package database

import (
	"encoding/base64"
	"encoding/json"
	"github.com/satori/go.uuid"
)

type AuthenticationToken struct {
	Id        string
	Token     string
	User      User
	IpAddress string
}

// language=sql
var AuthTokenTable = `
CREATE TABLE "auth_token" (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    token text NOT NULL,
    ip_address text NOT NULL,
    user_id uuid NOT NULL REFERENCES "user"(id)
)`

type Token struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

func CreateAuthenticationToken(user *User, ipAddress string) (string, error) {
	randomUuid := uuid.NewV4()
	token := Token{
		UserId: user.Id,
		Token:  randomUuid.String(),
	}

	db, err := Connect()
	if err != nil {
		return "", err
	}

	// language=sql
	_, err = db.Exec("INSERT INTO auth_token (token, user_id) VALUES ($1, $2)", token.Token, token.UserId)
	if err != nil {
		return "", err
	}

	return EncodeToken(token)
}

func EncodeToken(token Token) (string, error) {
	serializedToken, err := json.Marshal(token)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(serializedToken), nil
}

func DecodeToken(token string) (*Token, error) {
	result, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	deserializedToken := new(Token)
	err = json.Unmarshal(result, deserializedToken)

	return deserializedToken, err
}

func ValidateAuthenticationToken(token string) bool {
	db, err := Connect()
	if err != nil {
		return false
	}

	defer db.Close()

	count := new(int)

	deserializedToken, err := DecodeToken(token)
	if err != nil {
		return false
	}

	// language=sql
	err = db.Get(count, "SELECT COUNT(*) FROM auth_token WHERE token = $1 AND user_id = $2", deserializedToken.Token, deserializedToken.UserId)
	if err != nil || *count == 0 {
		return false
	}

	return true
}
