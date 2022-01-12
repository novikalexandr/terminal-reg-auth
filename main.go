package main

// TODO: Вынести логин и рег в отдельные пакеты

import (
	"database/sql"
	"fmt"
	"terminal-reg-auth/login"
	"terminal-reg-auth/reg"

	_ "github.com/mattn/go-sqlite3"
)

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
