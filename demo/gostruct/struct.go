package gostruct

import (
	"fmt"
	"math"
)

// ----------struct 结构体初始化-----------------

// Person 结构体
type Person struct {
	name string
	age  int
}

// Older 返回更大人结构体与年龄差
func Older(p1, p2 Person) (Person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

// TestStruct 测试结构体
func TestStruct() {

	var tom Person

	// 赋值初始化
	tom.name, tom.age = "Tom", 22

	// 两个字段都写清楚的初始化
	bob := Person{name: "Bod", age: 29}

	// 按照 struct 定义顺序初始化
	pau := Person{"Paul", 45}

	// Older tom 和 bob 相比谁更大
	tbOlder, tbDiff := Older(tom, bob)
	// Older tom 和 paul 相比谁更大
	tpOlder, tpDiff := Older(tom, pau)
	// Older bob 和 paul 相比谁更大
	bOlder, bpDiff := Older(bob, pau)

	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, bob.name, tbOlder.name, tbDiff)
	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, pau.name, tpOlder.name, tpDiff)
	fmt.Printf("Of %s and %s, %s is older by %d years\n", bob.name, pau.name, bOlder.name, bpDiff)

}

// --------结构体对象-----------

// Rectangle 计算长方形面积
type Rectangle struct {
	width, height float64
}

// Circle 计算圆形面积
type Circle struct {
	radius float64
}

// 计算长方形面积(面积只是图形对象的一个属性)
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// 计算原型面积
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

// CalcArea 计算图形面积
func CalcArea() {

	r1 := Rectangle{width: 10, height: 5}

	r2 := Rectangle{18, 2}

	c1 := Circle{radius: 5}

	c2 := Circle{6}

	fmt.Println("r1 area is :", r1.area())
	fmt.Println("r2 area is :", r2.area())
	fmt.Println("c1 area is :", c1.area())
	fmt.Println("c2 ares is :", c2.area())
}

// ------------------------------
// WHITE 常量
const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

// Color 自定义method
type Color byte

// Box 对象
type Box struct {
	width, height, depth float64
	color                Color
}

// BoxList 对象切片
type BoxList []Box // a slice of boxes

// Volume 对象面积
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

// SetColor 设定对象颜色
func (b *Box) SetColor(c Color) {
	b.color = c
}

// BiggestColor 最大对象的颜色
func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

// PanitItBlack 给对象上色
func (bl BoxList) PanitItBlack() {
	for i := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

// TestBox 测试盒子对象入口
func TestBox() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")
	boxes.PanitItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())

	fmt.Println("Obviously,now,the biggest one is", boxes.BiggestColor().String())
}

//----------method 继承--------------

// Human 结构体
type Human struct {
	name  string
	age   int
	phone string
}

// Student 结构体
type Student struct {
	Human  // 匿名字段
	school string
}

// Employee 结构体
type Employee struct {
	Human
	company string
}

// SayHi 在Human结构体上定义一个method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// SayHi 在Employee 结构体上定义一个method
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

// ExtendMethod 方法继承
func ExtendMethod() {

	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}

	sam := Employee{Human{"Sam", 45, "111-888-xxxx"}, "Golang Inc"}

	mark.SayHi()

	sam.SayHi()
}
