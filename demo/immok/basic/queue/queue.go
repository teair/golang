package queue

// Queue An FIFO queue
type Queue []interface{}

// Push the element into the queue.
// 	e.g. q.Push(123)
func (q *Queue) Push(i interface{}) {
	*q = append(*q, i)
	//*q = append(*q, i.(int))
}

// Pop element from head.
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]

	//return head
	// 如果想限定返回类型
	return head
	//return head.(int)
}

// IsEmpty Returns if the Queue is Empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
