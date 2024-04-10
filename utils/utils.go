package utils

import (
	"database/sql"
	"fmt"

	"github.com/akshay98322/Go-gorilla-mux-my-sql/constants"
	models "github.com/akshay98322/Go-gorilla-mux-my-sql/model"
)

func CreateTable(query string) {
	db, err := sql.Open(constants.DbDriver, constants.DbUser+":"+constants.DbPass+"@/"+constants.DbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Table created successfully..")
	}

}

func CreateUser(db *sql.DB, name, email string) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err := db.Exec(query, name, email)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(db *sql.DB, id int) (*models.User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := db.QueryRow(query, id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(db *sql.DB, id int, name, email string) error {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err := db.Exec(query, name, email, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(db *sql.DB, id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
