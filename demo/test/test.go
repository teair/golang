package test

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("下班喽~~~")
			return
		default:
			fmt.Println("认真摸鱼中，请勿打扰...")
		}
		time.Sleep(1 * time.Second)
	}
}

func TestMain() {
	ctx,cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		worker(ctx)
	}()
	time.Sleep(3 * time.Second)	// 工作3秒
	cancel()					// 3秒后发出停止指令
	wg.Wait()
}
