package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(name string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("name: %s, message %d", name, i)
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

func main() {
	c1 := msgGen("chan1")
	c2 := msgGen("chan2")
	c3 := msgGen("chan3")

	// 并行接收之goroutine+for
	f := fanIn(c1, c2, c3)
	// 并行接收之select
	//f := fanInBySelect(c1,c2)
	for {
		fmt.Println(<-f)
	}
}
