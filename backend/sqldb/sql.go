package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:txc5669857@tcp(127.0.0.1:3306)/gintest")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//新增
	/*res,err:=db.Exec("insert into `user`(`id`,`name`,`telephone`,`password`)values(0,'杨斌4','17377792524','123123')")
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		fmt.Println("insert success")
	}*/
	//删除
	/*del,err:=db.Exec("delete from `user`where `telephone`=?","17377792524")

	if err != nil {
		log.Fatal(err)
	}
	if del != nil {
		fmt.Println("delete the one who`s telephone is 17377792524")
	}*/

	//更新

	update, err := db.Query("update `user` set `password`=? where `telephone`=?", "123456", "17377792521")
	if err != nil {
		log.Fatal(err)
	}
	if update != nil {
		fmt.Println("update success")
	}

	//获取数据
	rows, err := db.Query("select * from user ")

	//查询
	for rows.Next() {
		var (
			id   int
			name string

			telephone string
			password  string
		)

		err = rows.Scan(&id, &name, &telephone, &password)
		fmt.Println(id, name, telephone, password)
		if err != nil {
			log.Fatal(err)
		}
	}

}
