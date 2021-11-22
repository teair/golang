package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done func()
}

func doWork(id int, w worker) {
	go func() {
		for n := range w.in {
			/*n,ok := <-c
			if !ok {
				// 如果channel为空则停止接收
				break
			}*/
			// 由于fmt涉及到一些IO操作所以goroutine会进行调度所以会显示无序
			fmt.Printf("worker %d, received:%c\n", id, n)
			w.done()
		}
	}()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func channelDemo() {

	var wg sync.WaitGroup
	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)

	for i, worker := range workers {
		worker.in <- 'a' + i // 给 channel 发送值
		//<-worker.done	// 由于是按顺序执行与并发理念不符
	}

	for m, worker := range workers {
		worker.in <- 'A' + m
		//<-worker.done	// 由于是按顺序执行与并发理念不符
	}

	wg.Wait()
}

func main() {
	channelDemo()
}
