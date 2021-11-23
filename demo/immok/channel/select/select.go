package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond) // 随机睡1500毫秒
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	go func() {
		for n := range c {
			/*n,ok := <-c
			if !ok {
				// 如果channel为空则停止接收
				break
			}*/
			time.Sleep(time.Second)
			// 由于fmt涉及到一些IO操作所以goroutine会进行调度所以会显示无序
			fmt.Printf("worker %d, received:%d\n", id, n)
		}
	}()
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	var values []int
	hasValue := false

	tm := time.After(10 * time.Second)

	tick := time.Tick(time.Second) // 每秒获取一次队列长度

	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		if hasValue {
			activeWorker = worker
		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout!") // 每次从channel获取的时间超过800毫秒就打印timeout
		case <-tick:
			fmt.Println("queue len = ", len(values))
		case <-tm:
			fmt.Println("byte") // 总的运行时间设为10秒
			return
		}
	}
}
