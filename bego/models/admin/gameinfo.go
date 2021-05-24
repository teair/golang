package admin

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type GameInfo struct {
	Id            int       `orm:"auto"`
	Gid           string    `orm:"size(50)"`
	AppName       string    `orm:"size(128)"`
	AppRename     string    `orm:"size(100)"`
	MarkOnline    int8      `orm:"size(5)"`
	GameSize      string    `orm:"size(32)"`
	Source        string    `orm:"size(10)"`
	Publicity     string    `orm:"size(150)"`
	Fuli          string    `orm:"size(255)"`
	GameIntroduce string    `orm:"type(text)"`
	OnlineTime    time.Time `orm:"auto_now;type(datetime)"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	// 连接数据库信息
}

// 可创建分发游戏列表
func GameList() {

}

// 获取分发游戏信息
func GetMsg(where interface{}) {
	/*if where != nil {
		o := orm.NewOrm()
		var fy_game_info []FyGameInfo
		num, _ := o.QueryTable("fy_game_info").All(&fy_game_info)
		log.Println("查询出的数据:",fy_game_info)
		log.Println("总记录数:",num)
	}*/
}

// 获取所有游戏列表
func GameInfoSelect() []GameInfo {
	o := orm.NewOrm()
	var fy_game_info []GameInfo
	num, err := o.QueryTable("表名").All(&fy_game_info)
	if num == 0 || err != nil {
		return nil
	}
	return fy_game_info
}
