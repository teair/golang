package adminmodel

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type GameLabel struct {
	Id         int       `orm:"int(11)"`
	Gid        string    `orm:"string(32)"`
	Type       string    `orm:"string(50)"`
	LabelId    string    `orm:"string(10)"`
	LabelName  string    `orm:"string(100)"`
	CreateTime time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModelWithPrefix("fy_", new(GameLabel))
}

// 判断当前游戏是否存在标签
func CurLabel(gid, gametype string) (g []GameLabel) {
	o := orm.NewOrm()
	labelcond := orm.NewCondition().And("Gid", gid).And("Type", gametype)
	o.QueryTable("fy_game_label").SetCond(labelcond).All(&g)
	return g
}
