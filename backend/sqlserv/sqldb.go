package sqlserv

import (
	"database/sql"
	"fmt"
	"ginDemo/backend/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func Init() {
	db, err := sql.Open("mysql", "root:txc5669857@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}
func CheckLogin(user model.User) bool {

	flag := false

	row, err := DB.Query("select `telephone`,`password` from user where `telephone`=? and `password`=?", user.Telephone, user.Password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(row)

	if row != nil {
		flag = true
	}
	DB.Close()
	return flag

}

func GetDbData(telephone string) []model.UserMsg {
	rows, err := DB.Query("SELECT username,plcontext,replycontext,rusername FROM \n(SELECT u.userid,u.username,u.telephone  FROM `user` u WHERE telephone=?)\n as a\nJOIN pl p\nON p.useridr = a.userid\nright JOIN reply r\non r.plid=p.plid\nleft JOIN (SELECT u.userid,(u.username)as rusername FROM `user` u) as u\nON u.userid= r.wuserid ", telephone)
	if err != nil {
		log.Fatal(err)
	}
	//查询

	userMsg := make([]model.UserMsg, 1)
	i := 0
	for rows.Next() {
		user := model.UserMsg{}
		fmt.Println(rows.Columns())
		err := rows.Scan(&user.Username, &user.Plcontext, &user.Replycontext, &user.Rusername)
		if err != nil {
			log.Fatal(err)
		}
		userMsg[i] = user
		fmt.Println(user)
		i++
	}
	DB.Close()

	fmt.Println(userMsg)

	return userMsg
}
