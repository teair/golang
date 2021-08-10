package main

import (
	"bego/models"
	_ "bego/routers"
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	var err error
	dsn, _ := beego.AppConfig.String("gorm::data_source")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "fy_", // 表名前缀，`User`表为`fy_users`
			SingularTable: true,  // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	if err != nil {
		panic(err)
	}
	models.Db = db
}

func main() {
	beego.Run()
}
