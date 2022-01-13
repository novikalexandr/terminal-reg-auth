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
		return fmt.Errorf("\nThe login field cannot be empty. Try entering the login again.\n")

	}

	var insertedPassword string
	fmt.Printf("Enter your password\n> ")
	fmt.Scanf("%s\n", &insertedPassword)
	switch insertedPassword {
	case "":
		fmt.Printf("\nThe password field cannot be empty. Try entering the password again.\n")
		Login(db)

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

	// Check if user whith this credintials exists.
	rows, err := db.Query("SELECT username, password FROM users WHERE username = ? AND password = ?;", insertedLogin, insertedPassword)
	if err != nil {
		return err
	}
	defer rows.Close()
	user := user{}
	for rows.Next() {
		err := rows.Scan(&user.username, &user.password)
		if err != nil {
			return err
		}
		fmt.Printf("Hi: %s\nYour password: %s\nYou have successfully logged in.\n", user.username, user.password)
	}

	return nil
}
