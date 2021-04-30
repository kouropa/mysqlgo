package main

import (
	"fmt"
	//"os"
	"database/sql"

	//_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// var (
// 	db  *gorm.DB
// 	err error
// )

//User is struct
// type User struct {
// 	id    int
// 	name  string
// 	email string
// }

// ...略

func main() {

	_, err := sql.Open("mysql", "root:rootpassword@tcp(dockerMySQL:3306)/sample_db")

	//_, err := gorm.Open("mysql", connection)

	if err != nil {
		panic(err)
	}
	if err == nil {
		fmt.Println("yessssss")
	}

	// rows, err := db.Query("SELECT * FROM users")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var user User
	// 	err := rows.Scan(&user.id, &user.name, &user.email)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	fmt.Println(user.id, user.name, user.email)
	// }
}
