package login

import (
	"database/sql"
	"fmt"
)

type user struct {
	username string
	password string
}

func Login(db *sql.DB) error {
	var insertedLogin string
	fmt.Printf("Enter your login\n> ")
	fmt.Scanf("%s\n", &insertedLogin)
	if insertedLogin == "" {
		return fmt.Errorf("login field cannot be empty. Enter login again")
	}

	var insertedPassword string
	fmt.Printf("Enter your password\n> ")
	fmt.Scanf("%s\n", &insertedPassword)
	switch insertedPassword {
	case "":
		err := fmt.Errorf("password field cannot be empty. Enter password again")
		return err
	}

	// Connect to db.
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	defer db.Close()

	user := user{}
	// Check if user whith this credintials exists.
	row := db.QueryRow("SELECT username, password FROM users WHERE username = ? AND password = ?;", insertedLogin, insertedPassword)
	err = row.Scan(&user.username, &user.password)
	if err != nil {
		return err //TODO: Корректно возвращать ошибку
	}

	fmt.Printf("Hi: %s\nYour password: %s\nYou have successfully logged in.\n", user.username, user.password)

	return nil
}
