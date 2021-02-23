package gostruct

import (
	"fmt"
	"math"
)

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
