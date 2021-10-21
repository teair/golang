package entry

import (
	"demo/immok/basic/queue"
	"fmt"
)

func QueueMain() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	q.Push(4)

	pop := q.Pop()
	fmt.Println("pop num is:", pop)
	pop = q.Pop()
	fmt.Println("pop num is:", pop)

	fmt.Println("the Queue IsEmpty:", q.IsEmpty())

	fmt.Println("the Queue is:", q)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	q.Push("abc") // interface 表示可以接受任何类型
	fmt.Println(q.Pop())

}
