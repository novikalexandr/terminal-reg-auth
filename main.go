package main

// TODO: Вынести логин и рег в отдельные пакеты

import (
	"database/sql"
	"fmt"
	"login/login"
	"reg/reg"

	_ "github.com/mattn/go-sqlite3"
)

// type user struct {
// 	username string
// 	password string
// }

// func login(db *sql.DB) error {
// 	var insertedLogin string
// 	fmt.Printf("Enter your login\n> ")
// 	fmt.Scanf("%s\n", &insertedLogin)

// 	var insertedPassword string
// 	fmt.Printf("Enter your password\n> ")
// 	fmt.Scanf("%s\n", &insertedPassword)

// 	// Connect to db.
// 	db, err := sql.Open("sqlite3", "database.db")
// 	if err != nil {
// 		return err
// 	}
// 	err = db.Ping()
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	// Check if user whith this credintials exists.
// 	rows, err := db.Query("SELECT username, password FROM users WHERE username = ? AND password = ?;", insertedLogin, insertedPassword)
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()
// 	user := user{}
// 	for rows.Next() {
// 		err := rows.Scan(&user.username, &user.password)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	fmt.Printf("Hi: %s\nYour password: %s\nYou have successfully logged in.\n", user.username, user.password)

// 	return nil
// }

// func reg(db *sql.DB) error {
// 	var login string
// 	fmt.Printf("Add you login\n> ")
// 	fmt.Scanf("%s\n", &login)

// 	var pass string
// 	fmt.Printf("Add you password\n> ")
// 	fmt.Scanf("%s\n", &pass)

// 	_, err := db.Exec("insert into users (username, password, reg_time) values ($1, $2, $3)",
// 		login, pass, time.Now())
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("Hi, %s your registration complete\n", login)
// 	return nil
// }

func main() {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		fmt.Printf("error while connecting to DB: %s\n", err)
		return
	}
	defer db.Close()

	var choose string
	fmt.Printf("Hi Men, do you want login or register?\n> ")
	fmt.Scanf("%s\n", &choose)
	switch choose {
	case "login":
		for {
			err := login.Login(db)
			if err != nil {
				fmt.Println(err)
				continue
			}
			return
		}

	case "register":
		for {
			err := reg.Reg(db)
			if err != nil {
				fmt.Println(err)
				continue
			}
			return
		}
	case "exit":
		return
	default:
		fmt.Println("Wrong choose!")
	}
}
