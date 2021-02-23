package time

import (
	"fmt"
	"time"
)

// Time 测试time包
func Time() {
	now := time.Now() // 获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 小时
	minute := now.Minute() // 分钟
	second := now.Second() // 秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// UnixTime 时间戳
func UnixTime() {
	now := time.Now()            // 获取当前时间
	timestamp1 := now.Unix()     // 时间戳
	timestamp2 := now.UnixNano() // 纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

// UnixToTime Unix to time
func UnixToTime(timestamp int64) {

	timeObj := time.Unix(timestamp, 0) // 将时间戳转为时间格式
	fmt.Println(timeObj)

	year := timeObj.Year()     // 年
	month := timeObj.Month()   // 月
	day := timeObj.Day()       // 日
	hour := timeObj.Hour()     // 小时
	minute := timeObj.Minute() // 分
	second := timeObj.Second() // 秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)

}

// Duration
func Duration() {

}
