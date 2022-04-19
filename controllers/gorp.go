package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbmap = initDb()

func initDb() *gorm.DB {

	dsn := "root:root@tcp(127.0.0.1:3306)/laravel?charset=utf8mb4&parseTime=True&loc=Local"
	dbmap, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	checkErr(err, "failed to connect database")
	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
