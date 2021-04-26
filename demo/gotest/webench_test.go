package gotest

import "testing"

func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {

	b.StopTimer()

	// 在这里可以做一些初始化的参数
	// 并且不影响测试函数的性能

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}
