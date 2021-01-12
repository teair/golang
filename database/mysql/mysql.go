package mysql

import (
	"fmt"
	// mysql 库
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Test 结构体是表tests的映射
type Test struct {
	ID      int    `db:"id"`
	Uname   string `db:"uname"`
	Company string `db:"company"`
	Duty    string `db:"duty"`
	Mobile  string `db:"mobile"`
}

// Db 数据库对象
var Db *sqlx.DB

// Mysql 测试
func Mysql() {

	// database 连接数据库
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/tests")

	if err != nil {
		fmt.Println("failed connect mysql:", err)
		return
	}

	// Db 数据库对象
	Db = database

	defer Db.Close()

	// 新增

	// i, err := Db.Exec("insert into test values(?,?,?,?,?)", nil, "小绿", "针织科技", "产品", "116")

	// if err != nil {
	// 	fmt.Println("insert failed:", err)
	// 	return
	// }

	// id, err := i.LastInsertId()

	// if err != nil {
	// 	fmt.Println("exec failed:", err)
	// 	return
	// }

	// fmt.Println("exec succ:", id)

	// 修改

	// u, err := Db.Exec("update test set mobile = ? where id=?", 176, 5)

	// if err != nil {
	// 	fmt.Println("update failed:", err)
	// 	return
	// }

	// row, err := u.RowsAffected()

	// fmt.Println("affected row:", row)

	// 删除

	// d, err := Db.Exec("delete from test where id = ?", 6)

	// if err != nil {
	// 	fmt.Println("delete failed:", err)
	// 	return
	// }

	// row, err := d.RowsAffected()

	// fmt.Println("affected row:", row)

	// 查看
	var test []Test

	e3 := Db.Select(&test, "select id,uname,company,duty,mobile from test")

	if e3 != nil {
		fmt.Println("read failed", e3)
		return
	}

	for i, v := range test {
		fmt.Println(i, v)
	}

	// 事务
	// conn, err := Db.Begin()

	// f1, e1 := Db.Exec("insert into test values(?,?,?,?,?)", nil, "张飞", "百度", "保安", "130")

	// if e1 != nil {
	// 	conn.Rollback()
	// 	fmt.Println("f1 failed,", e1)
	// 	return
	// }

	// id1, e1 := f1.LastInsertId()

	// if e1 != nil {
	// 	fmt.Println("get id1 failed:", e1)
	// 	return
	// }

	// fmt.Println("id1 is:", id1)

	// f2, e2 := Db.Exec("insert into test values(?,?,?,?,?)", nil, "刘备", "百度", "经理", "131")

	// if e2 != nil {
	// 	fmt.Println("f2 failed,", e2)
	// 	conn.Rollback()
	// 	return
	// }

	// id2, e2 := f2.LastInsertId()

	// if e2 != nil {
	// 	fmt.Println("get id2 failed:", e2)
	// 	return
	// }

	// fmt.Println("id2 is :", id2)

	// conn.Commit()

	// fmt.Println("insert success")

}
