package queue

// Queue 通过别名来扩展业务类型
type Queue []interface{}

// Push func (q *Queue) Push(i interface{}) {
// 如果想限定加入类型
//func (q *Queue) Push(i int) {
// 这种写法编译可以运行时报错
func (q *Queue) Push(i interface{}) {
	*q = append(*q, i)
	//*q = append(*q, i.(int))
}

// Pop func (q *Queue) Pop() interface{} {
//func (q *Queue) Pop() int {
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]

	//return head
	// 如果想限定返回类型
	return head
	//return head.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
