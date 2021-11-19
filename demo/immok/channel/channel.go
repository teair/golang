package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	go func() {
		for n := range c {
			/*n,ok := <-c
			if !ok {
				// 如果channel为空则停止接收
				break
			}*/
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

func channelDemo() {

	for i := 0; i < 10; i++ {
		createWorker(i) <- 'a' + i // 给 channel 发送值
	}

	for m := 0; m < 10; m++ {
		createWorker(m) <- 'A' + m
	}

	time.Sleep(time.Millisecond)
}

func bufferChannel() {
	c := make(chan int, 3) // 增加缓冲区可提高性能
	go worker(0, c)
	c <- 'A'
	c <- 'B'
	c <- 'C'
	c <- 'D'
	time.Sleep(time.Millisecond)
}

func closeChannel() {
	c := make(chan int, 3) // 增加缓冲区可提高性能
	go worker(0, c)
	c <- 'A'
	c <- 'B'
	c <- 'C'
	c <- 'D'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	channelDemo()
	bufferChannel()
	closeChannel()
}
