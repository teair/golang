package mysql

import (
	"fmt"

	"github.com/astaxie/beego/orm"

	// mysql 驱动
	_ "github.com/go-sql-driver/mysql"
)

// profile 用户表
type profile struct {
	ID                           int
	Uname, Company, Duty, Mobile string
}

func init() {
	// 注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 别名	数据库类型	连接参数	最大数据库连接
	orm.RegisterDataBase("default", "mysql", "", 30)
	// 注册定义的model(可同时注册多个model,相当于创建多个表的映射)
	orm.RegisterModel(new(profile))
	// 不存在则创建table
	// orm.RunSyncdb("default", false, true)
}

// OrmMain orm框架入口
func OrmMain() {

	// 打印调试
	orm.Debug = true

	// 导入必须的package之后，我们需要打开到数据库的链接，然后创建一个beego orm对象
	ormObj := orm.NewOrm()

	// 插入表
	// user := profile{Uname: "百里守约", Company: "阿里", Duty: "p4", Mobile: "012"}
	// _, err := ormObj.Insert(&user)
	// fmt.Println(user)

	// 更新表
	// user.Uname = "虞姬"
	// user.Company = "鹅厂"
	// user.Duty = "架构师"
	// user.Mobile = "255"
	// num, err := ormObj.Update(&user)
	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// 读取 one
	u := profile{ID: 2}
	err := ormObj.Read(&u)
	if err == orm.ErrNoRows {
		fmt.Println("不存在该记录!")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键!")
	} else {
		fmt.Println(u.ID, u.Uname, u.Company, u.Duty, u.Mobile)
	}

	// ReadOrCreate 尝试从数据库读取，不存在则创建一个
	// p := profile{Uname: "钟馗", Company: "驱魔集团", Duty: "CFO", Mobile: "110"}
	// if isCreated, id, err := ormObj.ReadOrCreate(&p, "Company"); err == nil {
	// 	if isCreated {
	// 		fmt.Println("New insert an Object. Id:", id)
	// 	} else {
	// 		fmt.Println("Get an object. Id:", id)
	// 	}
	// }

	// InsertMulti 同时插入多个对象
	// p := []profile{
	// 	{Uname: "马云", Company: "阿里巴巴", Duty: "执行董事", Mobile: "456"},
	// 	{Uname: "马化腾", Company: "腾讯", Duty: "CEO", Mobile: "789"},
	// 	{Uname: "张磊", Company: "高瓴资本", Duty: "创始人兼首席执行官", Mobile: "159"},
	// }
	// addNum, err := ormObj.InsertMulti(1, p)

	// if err == nil {
	// 	fmt.Printf("成功插入 %v 条数据", addNum)
	// }

	// Update 更新
	// p := profile{ID: 3}
	// // Read 返回的是error类型，就是说成功读到该数据时
	// if ormObj.Read(&p) == nil {
	// 	p.Uname = "申通通"
	// 	if num, err := ormObj.Update(&p); err == nil {
	// 		fmt.Println(num)
	// 	}
	// }

	//  Delete 删除
	// if num, err := ormObj.Delete(&profile{ID: 10}); err == nil {
	// 	fmt.Println(num)
	// }

}
