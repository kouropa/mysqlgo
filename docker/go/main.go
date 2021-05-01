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


type User struct {
	id    int
	name  string
	email string
}

// ...略

func main() {

	
	
	db, err := sql.Open("mysql", "kouropa:password@tcp(dockerMySQL:3306)/golang_db")//rootだとアクセスできなかった。あとホスト名はcomposeのコンテナネーム

	//_, err := gorm.Open("mysql", "root:rootpassword@tcp(dockerMySQL:3306)/sample_db")

	if err != nil {
		panic(err)
	}
	if err == nil {
		fmt.Println("yessssss")
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.id, &user.name, &user.email)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.id, user.name, user.email)
	}
}
