package main

import (
	"newsapi/config"
	"newsapi/factory"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConn := config.InitDB()
	r := gin.Default()

	factory.InitFactory(r, dbConn)
	r.Run()
}
