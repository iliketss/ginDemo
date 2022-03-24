package main

import (
	"ginDemo/backend/router"
	"ginDemo/backend/sqlserv"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	sqlserv.Init()

	router.Login(r)
	router.GetMessage(r)

	r.Run()
}
