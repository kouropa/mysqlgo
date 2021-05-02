package main

import (
	//"database/sql"
	"database/sql"
	"fmt"

	//"os"
	"github.com/jinzhu/gorm"

	//_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// var (
// 	db  *gorm.DB
// 	err error
// )
const (
	tableNameUser = "users"
)

var Db *sql.DB

// type User struct {
// 	id       int
// 	uuid     string
// 	name     string
// 	email    string
// 	password string
// }

// ...略

func gormConnect() *gorm.DB {
	// DBMS := "mysql"
	// USER := "kouropa"
	// PASS := "password"
	// PROTOCOL := tcp(dockerMySQL:3306)
	// DBNAME := "dockerMy"
	// CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	// db, err := gorm.Open(DBMS, CONNECT)
	db, err := gorm.Open("mysql", "kouropa:password@tcp(dockerMySQL:3306)/golang_db") //rootだとアクセスできなかった。あとホスト名はcomposeのコンテナネーム

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected: ", &db)
	return db
}

func main() {

	//Db, err := gorm.Open("mysql", "kouropa:password@tcp(dockerMySQL:3306)/golang_db") //rootだとアクセスできなかった。あとホスト名はcomposeのコンテナネーム

	//_, err := gorm.Open("mysql", "root:rootpassword@tcp(dockerMySQL:3306)/sample_db")
	db := gormConnect()
	defer db.Close()

}
