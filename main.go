package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func login() error {
	var insertedLogin string
	fmt.Printf("Enter your login\n> ")
	fmt.Scanf("%s\n", &insertedLogin)

	var insertedPassword string
	fmt.Printf("Enter your password\n> ")
	fmt.Scanf("%s\n", &insertedPassword)

	passwordBytes, err := os.ReadFile(insertedLogin + ".txt")
	if err != nil {
		return fmt.Errorf("validation error")
	}
	passwordFromFile := string(passwordBytes)
	if insertedPassword != passwordFromFile {
		return fmt.Errorf("validation error: password unexist")
	}
	return nil
}

func reg(db *sql.DB) error {
	var login string
	fmt.Printf("Add you login\n> ")
	fmt.Scanf("%s\n", &login)

	var pass string
	fmt.Printf("Add you password\n> ")
	fmt.Scanf("%s\n", &pass)

	_, err := db.Exec("insert into users (username, password, reg_time) values ($1, $2, $3)",
		login, pass, time.Now())
	if err != nil {
		panic(err)
	}

	// file, err := os.Create(login + ".txt")
	// if err != nil {
	// 	return fmt.Errorf("Unable to create file:")
	// }

	// defer file.Close()
	// file.WriteString(pass)

	fmt.Printf("Hi, %s your registration complete\n", login)
	return nil
}

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
			err := login()
			if err != nil {
				fmt.Println(err)
				continue
			}
			return
		}

	case "register":
		for {
			err := reg(db)
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
