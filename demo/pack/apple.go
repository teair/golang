package pack

import(
	. "fmt"
)

type Weekday int

const(
	Sunday Weekday = iota		// 常量生成器
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Psd struct {
	hobby string
	Mobile string
}

func Apple() {
	Println("Apple!")
	//return ps.hobby,ps.Mobile
	//Println(Sunday,Monday,Tuesday,Wednesday,Thursday,Friday,Saturday)
}
