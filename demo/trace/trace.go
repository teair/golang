package trace

import (
	"fmt"
	"os"
	"runtime/trace"
)

func TraceMain() {

	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	trace.Stop()

	/*go func() {
		fmt.Println("1")
	}()

	go func() {
		fmt.Println("2")
	}()

	go func() {
		fmt.Println("3")
	}()*/

	go func() {
		fmt.Println("1")
		go func() {
			fmt.Println("2")
			go func() {
				fmt.Println("3")
			}()
		}()
	}()

	fmt.Println("hello world!")
}
