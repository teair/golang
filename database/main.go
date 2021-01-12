package main

import (
	"database/database/mysql"
	"database/database/redis"
	"fmt"
)

func main() {
	fmt.Println("11111")
	redis.Test()
	mysql.Mysql()
}
