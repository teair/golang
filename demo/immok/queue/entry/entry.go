package entry

import (
	"demo/immok/queue"
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
}
