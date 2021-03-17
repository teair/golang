package mysql

import (
	"fmt"

	"github.com/astaxie/beego/orm"

	// mysql 驱动
	_ "github.com/go-sql-driver/mysql"
)

// User 表
type User struct {
	Id, Age, Sex int
	Address      string
	Profile      *Profile `orm:"rel(one)"`      // OneToOne relation
	Post         []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系

}

// Profile 用户表
type Profile struct {
	Id                           int
	Uname, Company, Duty, Mobile string
	User                         *User `orm:"reverse(one)"` // OneToOne relation(设置一对一反向关系)
}

type Tag struct {
	Id   int
	Name string
	Post []*Post `orm:"reverse(many)"` // 设置多对多反向关系
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"` // 设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

func init() {
	// 注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 别名	数据库类型	连接参数	最大数据库连接
	orm.RegisterDataBase("default", "mysql", "", 30)
	// 注册定义的model(可同时注册多个model,相当于创建多个表的映射)
	orm.RegisterModel(new(Profile), new(User), new(Post), new(Tag))
	// 不存在则创建table
	// orm.RunSyncdb("default", false, true)
}

// OrmMain orm框架入口
func OrmMain() {

	// 开启sql打印调试
	orm.Debug = true

	// 导入必须的package之后，我们需要打开到数据库的链接，然后创建一个beego orm对象
	ormObj := orm.NewOrm()

	// 插入表
	// profile := Profile{Uname: "艾克", Company: "时空穿梭者", Duty: "Magic", Mobile: "266"}
	// _, err := ormObj.Insert(&profile)
	// fmt.Println(profile)

	// 关联插入单条记录
	// profile := new(Profile)
	// user := new(User)
	// post := new(Post)
	// profile.Uname, profile.Company, profile.Duty, profile.Mobile, profile.User = "扎克伯格", "Facebook", "脸书创始人", "245", user
	// user.Age, user.Sex, user.Address, user.Profile = 55, 1, "美国", profile
	// post.Title, post.User = "简简单单的打工人", user
	// _, err := ormObj.Insert(profile)
	// if err == nil {
	// 	fmt.Println(profile)
	// }
	// _, err = ormObj.Insert(user)
	// if err == nil {
	// 	fmt.Println(user)
	// }
	// _, err = ormObj.Insert(post)
	// if err == nil {
	// 	fmt.Println(post)
	// }

	// User 和 Profile 是OneToOne关系
	// user := &User{Id: 2}
	// ormObj.Read(user)
	// if user.Profile != nil {
	// 	ormObj.Read(user.Profile)
	// }
	// fmt.Println(user, user.Profile)

	// 直接关联查询(所谓关联就是存储关联信息的指针)
	// user := &User{}
	// ormObj.QueryTable("user").Filter("Id", 3).RelatedSel().One(user)
	// // 自动查询到了Profiel
	// fmt.Println(user.Profile)
	// // 由于在Profile里定义了反向关系的User，所以 Profile 里的 User 也是自动赋值过的，可以直接调用
	// fmt.Println(user.Profile.User)

	// 通过User查询Profile
	// var profile Profile
	// err := ormObj.QueryTable("profile").Filter("User__Id", 2).One(&profile)
	// if err == nil {
	// 	fmt.Println(profile)
	// }

	// Post 和 User 是 ManyToOne 关系，也就是 ForeignKey 为 User
	// var posts []*Post
	// num, err := ormObj.QueryTable("post").Filter("User", 3).RelatedSel().All(&posts)
	// if err == nil {
	// 	fmt.Printf("%d posts read \n", num)
	// 	for _, post := range posts {
	// 		fmt.Printf("Id: %d, UserName: %s, Title:%s \n", post.Id, post.User.Address, post.Title)
	// 	}
	// }

	// 根据 Post.Title 查询对应的User
	var post []*Post
	_, err := ormObj.QueryTable("post").Filter("Title__contains", "教育").RelatedSel().All(&post)
	if err == nil {
		for _, v := range post {
			fmt.Printf("The Id is: %v, Name is %s, Title is: %s, Age is: %v, Address is: %s\n", v.Id, v.User.Profile.Uname, v.Title, v.User.Age, v.User.Address)

		}
	}

	// 更新表
	// user.Uname = "虞姬"
	// user.Company = "鹅厂"
	// user.Duty = "架构师"
	// user.Mobile = "255"
	// num, err := ormObj.Update(&user)
	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// 读取 one
	// u := Profile{ID: 2}
	// err := ormObj.Read(&u)
	// if err == orm.ErrNoRows {
	// 	fmt.Println("不存在该记录!")
	// } else if err == orm.ErrMissPK {
	// 	fmt.Println("找不到主键!")
	// } else {
	// 	fmt.Println(u.ID, u.Uname, u.Company, u.Duty, u.Mobile)
	// }

	// ReadOrCreate 尝试从数据库读取，不存在则创建一个
	// p := Profile{Uname: "钟馗", Company: "驱魔集团", Duty: "CFO", Mobile: "110"}
	// if isCreated, id, err := ormObj.ReadOrCreate(&p, "Company"); err == nil {
	// 	if isCreated {
	// 		fmt.Println("New insert an Object. Id:", id)
	// 	} else {
	// 		fmt.Println("Get an object. Id:", id)
	// 	}
	// }

	// InsertMulti 同时插入多个对象
	// p := []Profile{
	// 	{Uname: "马云", Company: "阿里巴巴", Duty: "执行董事", Mobile: "456"},
	// 	{Uname: "马化腾", Company: "腾讯", Duty: "CEO", Mobile: "789"},
	// 	{Uname: "张磊", Company: "高瓴资本", Duty: "创始人兼首席执行官", Mobile: "159"},
	// }
	// addNum, err := ormObj.InsertMulti(1, p)
	// if err == nil {
	// 	fmt.Printf("成功插入 %v 条数据", addNum)
	// }

	// Update 更新
	// p := Profile{ID: 3}
	// // Read 返回的是error类型，就是说成功读到该数据时
	// if ormObj.Read(&p) == nil {
	// 	p.Uname = "申通通"
	// 	if num, err := ormObj.Update(&p); err == nil {
	// 		fmt.Println(num)
	// 	}
	// }

	//  Delete 删除
	// if num, err := ormObj.Delete(&Profile{ID: 10}); err == nil {
	// 	fmt.Println(num)
	// }

	// 高级查询 orm 以 QuerySeter 来组织查询

	// 获取QuerySeter对象
	// var p Profile
	// qs := ormObj.QueryTable(p)
	// 也可以直接使用对象作为表名
	// pro := new(Profile)
	// qs := ormObj.QueryTable(pro)

	// 计算符合条件数量
	// cnt, err := qs.Filter("ID__in", 2, 24, 25, 26).Count() // select count(*) from profile;
	// if err == nil {
	// 	fmt.Printf("Count num: %v", cnt)
	// }

	// 判断数据是否存在
	// exist := ormObj.QueryTable(p).Filter("ID", 2).Exist()
	// fmt.Println(exist)

	// 原子操作加(orm.ColAdd)减(orm.ColMinus)乘(orm.ColMultiply)除(orm.ColExcept)字段值
	// num, err := ormObj.QueryTable(p).Update(orm.Params{
	// 	"mobile": orm.ColValue(orm.ColAdd, 10),
	// })
	// if err == nil {
	// 	fmt.Println("受影响行数:", num)
	// }

	// All 返回对应的结果集对象
	// var pro []*Profile
	// num, err := ormObj.QueryTable("profile").Filter("ID__gt", 2).All(&pro)
	// fmt.Printf("Returned Rows Num: %v,%v", num, err)

	// One 尝试返回单条记录
	// var pro Profile
	// err := ormObj.QueryTable("profile").Filter("uname", "马云").One(&pro)
	// if err == orm.ErrMultiRows {
	// 	// 多条的时候报错
	// 	fmt.Printf("Returned Multi Rows Not One!")
	// }
	// if err == orm.ErrNoRows {
	// 	// 没有找到记录
	// 	fmt.Printf("Not row found!")
	// }
	// fmt.Println(pro)

	// Values 返回结果集的 key => vlaue
	// var maps []orm.Params
	// num, err := ormObj.QueryTable("profile").Values(&maps)
	// if err == nil {
	// 	fmt.Printf("Result Nums: %d\n", num)
	// 	for _, m := range maps {
	// 		fmt.Println(m["ID"], m["Uname"]) // 返回指定的Field数据
	// 	}
	// }

	// // 可以指定 expr 返回指定的 Field
	// var lists []orm.ParamsList
	// num, err = ormObj.QueryTable("profile").ValuesList(&lists, "ID", "Duty")
	// if err == nil {
	// 	fmt.Println("Result nums: ", num)
	// 	for _, row := range lists {
	// 		fmt.Printf("ID: %d, Duty: %s\n", row[0], row[1])
	// 	}
	// }

	// // 只返回特定的 Field 值，将结果集展开到单个slice里
	// var strs orm.ParamsList
	// num, err = ormObj.QueryTable("profile").ValuesFlat(&strs, "Mobile")
	// if err == nil {
	// 	fmt.Printf("Results Nums: %d\n", num)
	// 	fmt.Printf("All Profile Names: %s", strs)
	// }

	// 关联查询 User 和 Profile 是 OneToOne 的关系
	// user := &User{ID: 3}
	// ormObj.Read(user)
	// if user.Profile != nil {
	// 	ormObj.Read(user.Profile)
	// }
}
