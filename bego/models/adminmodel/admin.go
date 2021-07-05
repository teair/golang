package adminmodel

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type SystemUser struct {
	Id         int       `orm:"auto"`
	Username   string    `orm:"size(32)"`
	Password   string    `orm:"size(32)"`
	CreateTime time.Time `orm:"type(datetime)"`
}

func init() {
	// 设置表前缀
	orm.RegisterDataBase("default", "mysql", "webgame:123456@tcp(182.92.149.107:3306)/webgame?charset=utf8")
	orm.RegisterModelWithPrefix("fy_", new(SystemUser))
}

// 判断账号密码是否正确
func CheckAdmin(username, password string) (bool, error) {

	o := orm.NewOrm()

	var sys_user SystemUser

	err := o.QueryTable("fy_system_user").Filter("username", username).One(&sys_user)

	if err != nil {
		log.Println("login model CheckAdmin failed:", err)
	}
	if sys_user.Username != "admin" {
		return false, errors.New("用户名错误!")
	}

	// 加密用户输入密码
	pwd := []byte(password)
	md5ctx := md5.New()
	md5ctx.Write(pwd)
	crypwd := fmt.Sprintf("%x", md5ctx.Sum(nil))

	if sys_user.Password != crypwd {
		return false, errors.New("密码错误!")
	}

	return true, nil
}
