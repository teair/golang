package web

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

var datas []string

func init() {
	// 对Block进行采集,通过调用该方法设置采集的频率,若不设置或设置小于等于0的数字也会认为关闭
	runtime.SetBlockProfileRate(1)
	// 对互斥锁(Mutex)进行采集,通过调用该方法设置采集频率,若不设置或没有设置大于0的数字默认是不进行采集的
	runtime.SetMutexProfileFraction(1)
}

func PprofMain() {
	go func() {
		for {
			log.Printf("len: %d", Add("go-programming-tour-book"))
			time.Sleep(time.Millisecond * 10)
		}
	}()
	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}

func Add(str string) int {
	data := []byte(str)
	datas = append(datas, string(data))
	return len(datas)
}

// Mutex Profiling(模拟阻塞情况)
// 1. go tool pprof http://localhost:6061/debug/pprof/mutex(go tool pprof http://localhost:6061/debug/pprof/block)
// 2. 运行 top 命令查看阻塞排名情况
// 3. list 函数名=>查看指定函数的代码情况
func MutexMain() {
	// 同步锁
	var m sync.Mutex
	var datas = make(map[int]struct{})
	for i := 0; i < 999; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			datas[i] = struct{}{}
		}()
	}
	_ = http.ListenAndServe("0.0.0.0:6061", nil)
}
