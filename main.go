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

	var userChoose string
	fmt.Printf("Hi Men, do you want login or register?\n> ")
	fmt.Scanf("%s\n", &userChoose)
	switch userChoose {
	case "login":
		for { //запуск бесконечного цикла
			err := login.Login(db) // err = результату отработки Login()
			if err != nil {        // если есть ошибка то мы выполняет то что нижу
				fmt.Println(err)
				continue // запускает цикл заново
			}
			return //убивает main() если нет ошибки
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
		fmt.Printf("Wrong choose: \"%s\", expect [login, register, exit].\n", userChoose)
	}
}
