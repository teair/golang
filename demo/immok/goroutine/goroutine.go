package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 1000; i++ {
		go func(i int) { // 如果不传ii会race condition【数据冲突】
			for {
				// 此处能继续向下运行是因为fmt包中做了一个IO的操作【其中有协程之间的切换】,所以会交出控制权给其他成员
				fmt.Printf("hello from goroutine %d\n", i)
				//s[ii]++
				//runtime.Gosched()
			}
		}(i)
	}

	time.Sleep(time.Minute) // 睡一毫秒
}
