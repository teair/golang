package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// OnCemat 结构体是表on_cemat的映射
type OnCemat struct {
	ID      uint8  `db:"id"`
	Uname   string `db:"uname"`
	Company string `db:"company"`
	Duty    string `db:"duty"`
	Mobile  string `db:"mobile"`
}

// Db 对象
var Db *sqlx.DB

func main() {

	// database 连接数据库
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/cemat")

	if err != nil {
		fmt.Println("failed connect mysql,", err)
		return
	}

	Db = database

	// defer Db.Close() // 注意，这行代码要写在上面err判断的下面
	defer Db.Close()

	// 新增
	// r, err := Db.Exec("insert into on_cemat(uname,company,duty,mobile) values(?,?,?,?)", "小红", "方游", "技术", "115")

	// if err != nil {
	// 	fmt.Println("exec failed,", err)
	// 	return
	// }

	// id, err := r.LastInsertId()

	// if err != nil {
	// 	fmt.Println("exec failed, ", err)
	// 	return
	// }

	// fmt.Println("insert succ,", id)

	// 查询
	var OnCemat []OnCemat

	err = Db.Select(&OnCemat, "select * from on_cemat where id = ?", 1)

	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}

	for _, v := range OnCemat {
		fmt.Println(v)
	}

	// 修改
	// res, err := Db.Exec("update on_cemat set mobile = ? where id = ?", 111, 8)

	// if err != nil {
	// 	fmt.Println("update failed,", err)
	// 	return
	// }

	// row, err := res.RowsAffected()

	// if err != nil {
	// 	fmt.Println("rows failed,", err)
	// 	return
	// }

	// fmt.Println("update succ:", row)

	// 删除
	// res, err := Db.Exec("delete from on_cemat where id = ?", 1)

	// if err != nil {
	// 	fmt.Println("failed delete:", err)
	// 	return
	// }

	// row, err := res.RowsAffected()

	// if err != nil {
	// 	fmt.Println("get affect failed:", err)
	// }

	// fmt.Println("delete succ:", row)

	// mysql 事务
	// conn, err := Db.Begin()

	// if err != nil {
	// 	fmt.Println("begin failed:", err)
	// 	return
	// }

	// f1, err := Db.Exec("insert into on_cemat values (?,?,?,?,?)", nil, "申通", "物联网", "技术", "211")

	// if err != nil {
	// 	fmt.Println("insert f1 failed:", err)
	// 	conn.Rollback()
	// 	return
	// }

	// f1id, err1 := f1.LastInsertId()

	// if err1 != nil {
	// 	fmt.Println("f1id get failed:", err1)
	// 	conn.Rollback()
	// 	return
	// }

	// fmt.Println("Insert f1id:", f1id)

	// f2, err2 := Db.Exec("insert into on_cemat values(?,?,?,?,?)", nil, "立新", "方游", "技术", "222")

	// if err2 != nil {
	// 	fmt.Println("f2 insert failed:", err2)
	// 	conn.Rollback()
	// 	return
	// }

	// f2id, err := f2.LastInsertId()

	// if err != nil {
	// 	fmt.Println("get f2id failed:", err)
	// 	conn.Rollback()
	// 	return
	// }

	// fmt.Println("Insert f2id:", f2id)

	// fmt.Println("All Insert successfull!")

	// conn.Commit()
}
