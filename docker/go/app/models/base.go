package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err := sql.Open("mysql", "kouropa:password@tcp(dockerMySQL:3306)/golang_db") //rootだとアクセスできなかった。あとホスト名はcomposeのコンテナネーム
	if err != nil {
		log.Fatalln(err)
	}
	if err == nil {
		fmt.Println("OK,connect")
		fmt.Println(Db)

	}

	// cmdU := fmt.Sprintln(`CREATE TABLE IF NOT EXISTS %s(
	// 	id INTEGER PRIMARY KEY AOTOINCRIMENT,
	// 	uuid STRING NOT NULL UNIQUE,
	// 	name STRING,
	// 	email STRING,
	// 	password STRING,
	// 	created_at DATETIME`, tableNameUser)
	// Db.Exec(cmdU)
}
