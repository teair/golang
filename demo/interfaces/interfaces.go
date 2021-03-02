package interfaces

import "fmt"

// Human 对象
type Human struct {
	name, phone string
	age         int
}

// Student 对象
type Student struct {
	Human
	school string
	loan   float32
}

// Employee 对象
type Employee struct {
	Human
	company string
	money   float32
}

// SayHi Human对象实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s , Call me on %s", h.name, h.phone)
}

// Sing Human对象实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("la la la la la la la...", lyrics)
}

// Guzzle Human对象实现Guzzle方法
func (h Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// SayHi Employee重载Human的SayHi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

// BorrowMoney Studnet实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount
}

// SpendSalary Employee实现SpendSalary方法
func (e Employee) SpendSalary(amount float32) {
	e.money -= amount
}

// 定义interface

// Men interface
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

// YoungChap interface
type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

// ElderGent interface
type ElderGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}
