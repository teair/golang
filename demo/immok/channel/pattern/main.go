package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(name string, done chan struct{}) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(5000)) * time.Millisecond):
				c <- fmt.Sprintf("name: %s, message %d", name, i)
			case <-done: // 任务中断与优雅退出(双向channel)
				fmt.Println("cleaning up!")
				time.Sleep(2 * time.Second) // 中断花费了两秒
				fmt.Println("cleanup done!")
				done <- struct{}{} // 通知调用程序清理完成
				return
			}
			i++
		}
	}()
	return c
}

// fanIn 同时接收两个通道的数据,避免等待耗费的时间
func fanIn(chs ...chan string) chan string {
	c := make(chan string)
	for _, ch := range chs {
		//chCopy := ch	// ch 只有一份,chCopy有两份
		// 通过函数传参拷贝来避免for range循环变量的坑
		go func(in chan string) {
			for {
				c <- <-in
			}
		}(ch)
	}
	return c
}

func fanInBySelect(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}
		}
	}()
	return c
}

// noneBlockingWait 非阻塞机制
func noneBlockingWait(c chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default:
		return "", false
	}
}

// timeoutWait 超时等待
func timeoutWait(c chan string, timeout time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m, true
	case <-time.After(timeout): // 超过该时间则返回空
		return "", false
	}
}

func main() {
	/*c1 := msgGen("chan1")
	c2 := msgGen("chan2")
	c3 := msgGen("chan3")

	// 并行接收之goroutine+for
	f := fanIn(c1, c2, c3)
	// 并行接收之select
	//f := fanInBySelect(c1,c2)
	for {
		fmt.Println(<-f)
	}*/

	done := make(chan struct{}) // 任务中断channel
	c := msgGen("chan", done)
	for i := 0; i < 5; i++ {
		if m, ok := timeoutWait(c, 2*time.Second); ok {
			fmt.Println("Received from chan:", m)
		} else {
			fmt.Println("timeout!")
		}
	}
	done <- struct{}{}      // 第一个{}为结构的定义,第二个{}是初始化这个结构
	time.Sleep(time.Second) // 给中断操作预留一秒钟的时间
	<-done                  // 优雅退出(双向通信)
}
