package reg

import (
	"database/sql"
	"fmt"
	"time"
)

func Reg(db *sql.DB) error {
	var login string
	fmt.Printf("Add you login\n> ")
	fmt.Scanf("%s\n", &login)

	var pass string
	fmt.Printf("Add you password\n> ")
	fmt.Scanf("%s\n", &pass)

	_, err := db.Exec("insert into users (username, password, reg_time) values ($1, $2, $3)", login, pass, time.Now())
	if err != nil {
		return err
	}

	fmt.Printf("Hi, %s your registration complete\n", login)
	return nil
}
