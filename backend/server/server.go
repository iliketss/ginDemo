package server

import (
	"fmt"
	"ginDemo/backend/model"
	"ginDemo/backend/sqlserv"
	"github.com/gin-gonic/gin"
	"log"
)

func LoginServer(c *gin.Context) {

	user := model.User{}

	err := c.BindJSON(&user)
	fmt.Println(user)
	if err != nil {
		log.Fatal(err)
	}

	if sqlserv.CheckLogin(user) {
		c.JSON(200, gin.H{
			"msg": "Login Success",
		})

	} else {

		c.JSON(400, gin.H{
			"msg": "用户名或者密码错误",
		})
	}

	//sqlserv.CheckLogin(user)

}
func GetMsg(c *gin.Context) {
	telephone := c.Query("telephone")
	fmt.Println(telephone)

	userMsg := sqlserv.GetDbData(telephone)

	c.JSON(200, gin.H{
		"data": userMsg,
	})

}
