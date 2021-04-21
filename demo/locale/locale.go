package locale

import (
	"encoding/json"
	"fmt"
	"time"
)

var locales map[string]map[string]string

type Data struct {
	Pea, Bean, HowOld, TimeZone, Money string
}

func LocaleMain() {

	locales = make(map[string]map[string]string, 2)

	// en
	en := make(map[string]string, 10)
	en["pea"] = "pea"
	en["bean"] = "bean"
	en["how_old"] = "I am %d years old!"
	en["time_zone"] = "America/Chicago"
	en["date_format"] = "2006-01-02 15:04:05" // 本地化时间格式
	en["money"] = "USD %d"
	locales["en"] = en

	// zh-CN
	cn := make(map[string]string, 10)
	cn["pea"] = "豌豆"
	cn["bean"] = "毛豆"
	cn["how_old"] = "我今年%d岁了!"
	cn["time_zone"] = "Asia/Shanghai"
	cn["date_format"] = "2006年-01月-02日 15时04分05秒" // 本地化时间格式
	cn["money"] = "￥%d元"
	locales["cn"] = cn

	lang := "en"

	loca, _ := time.LoadLocation(msg(lang, "time_zone")) // 获取并设定本地时区

	// 获取当前时区的当前时间
	t := time.Now().In(loca)

	/*fmt.Println(msg(lang,"pea"))
	fmt.Println(msg(lang,"bean"))
	fmt.Printf(msg(lang,"how_old")+"\n",30)
	fmt.Printf("当前时区：%v\n",t.Format("2006-01-02 15:04:05"))
	fmt.Println(date(msg(lang, "date_format"),t))*/

	d := &Data{
		Pea:      msg(lang, "pea"),
		Bean:     msg(lang, "bean"),
		HowOld:   fmt.Sprintf(msg(lang, "how_old"), 30),
		TimeZone: date(msg(lang, "date_format"), t),
		Money:    money_format(msg(lang, "money"), 18),
	}

	j, _ := json.Marshal(&d)

	fmt.Println(string(j))
}

func msg(locale, key string) string {
	if v, ok := locales[locale]; ok {
		if v2, ok2 := v[key]; ok2 {
			return v2
		}
	}
	return ""
}

// 日期格式化
func date(formate string, t time.Time) string {
	return t.Format(formate)
}

// 货币格式化
func money_format(format string, money int64) string {
	return fmt.Sprintf(format, money)
}
