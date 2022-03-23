package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	var name string
	db, err := sql.Open("mysql", "root:txc5669857@tcp(127.0.0.1:3306)/gintest")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//获取单行数据
	err = db.QueryRow("select name from user where id = ?", 1).Scan(&name)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}
