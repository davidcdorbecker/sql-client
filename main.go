package main

import (
	"errors"
	"fmt"
	"github.com/davidcdorbecker/sqlclient/sql_client"
)

const (
	userName = "root"
	password = "admin2684"
	host     = "127.0.0.1:3306"
	schema   = "test_users_db"
)

var (
	dbClient sql_client.SqlClientInterface
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func init() {
	var err error
	sql_client.StartMockServer()
	dbClient, err = sql_client.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", userName, password, host, schema))
	if err != nil {
		panic(err)
	}

	fmt.Println("connection success")
}

func main() {
	user, err := GetUser(123)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}

func GetUser(id int) (*User, error) {
	sql_client.AddMock(sql_client.Mock{
		Query:   "SELECT * FROM users WHERE id=?",
		Args:    []interface{}{123},
		Error:   nil,
		Columns: []string{"id", "name", "email"},
		Row:     []interface{}{1, "Test", "test@gmail.com"},
	})

	row := dbClient.QueryRow("SELECT * FROM users WHERE id=?", id)
	if row == nil {
		return nil, errors.New("user not found")
	}

	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
