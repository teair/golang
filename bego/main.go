package main

import (
	"bego/models"
	_ "bego/routers"
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	var err error
	dsn, _ := beego.AppConfig.String("gorm::data_source")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	models.DB = db
}

func main() {
	beego.Run()
}
