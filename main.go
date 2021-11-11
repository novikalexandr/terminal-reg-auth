package main

import (
	"fmt"
	"os"
)

func login() error {
	var insertedLogin string
	fmt.Printf("Enter your login\n> ")
	fmt.Scanf("%s\n", &insertedLogin)

	var insertedPassword string
	fmt.Printf("Enter your password\n> ")
	fmt.Scanf("%s\n", &insertedPassword)

	// Мы пытаемся найти и открыть файл, название которого равно введенному пользователем логину.
	loginFile, err := os.Open(insertedLogin + ".txt")
	if err != nil { // если не удалось найти и открыть файл, который запрашивает user.
		fmt.Println("Login unexist! Enter again.")
		return login()
	}
	defer loginFile.Close()

	// Читаем содержимое логин файла.
	passwordBytes := []byte{}
	_, err = loginFile.Read(passwordBytes)
	if err != nil {
		return fmt.Errorf("Password unexist!")
	}
	passwordFromFile := string(passwordBytes)

	if insertedPassword != passwordFromFile {
		return fmt.Errorf("Password unexist!")
	}
	return nil
}

func reg() error {
	var login string
	fmt.Printf("Add you login\n> ")
	fmt.Scanf("%s\n", &login)

	var pass string
	fmt.Println("Add you password\n> ")
	fmt.Scanf("%s\n", &pass)

	file, err := os.Create(login + ".txt")
	if err != nil {
		return fmt.Errorf("Unable to create file:")
	}

	defer file.Close()
	file.WriteString(pass)

	fmt.Println("Done!)")
	return nil
}

func main() {
	var choose string
	fmt.Printf("Hi Men, do you want login or register?\n> ")
	fmt.Scanf("%s\n", &choose)
	switch choose {
	case "login":
		err := login()
		if err != nil {
			fmt.Println(err)
			return
		}

	case "register":
		err := reg()
		if err != nil {
			fmt.Println(err)
			return
		}
	case "exit":
		return
	default:
		fmt.Println("Wrong choose!")
	}
}
