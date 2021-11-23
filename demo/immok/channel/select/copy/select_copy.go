package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	c := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			c <- i
			i++
		}
	}()
	return c
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
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

func main() {

	var c1, c2 = generator(), generator() // 初始化通道并给通道传值

	w := createWorker(0) // 返回一个只能写入的通道

	n := 0

	var values []int

	tm := time.After(10 * time.Second) // 10秒钟后会给返回的通道发送一个值

	tick := time.Tick(time.Second) // 每秒钟给返回的通道发送一个值

	for {
		// 这是一个临时通道,在无值时启用,有值时会替换为消耗输出的通道w
		var activeWorker chan<- int
		var activeValue int
		// 如果有值则把值传入w通道,没有值则等待其他case输出
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:] // 消耗一个更新下values队列
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout!") // 如果两次select间隔超过800毫秒则中这个case
		case <-tick:
			fmt.Println("queue len =", len(values))
		case <-tm:
			fmt.Println("bye!")
			return
		}
	}

}
