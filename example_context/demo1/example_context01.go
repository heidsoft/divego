package main

import (
	"context"
	"fmt"
	"time"
)

// https://golangbyexample.com/using-context-in-golang-complete-guide/
func main() {
	// 创建root context 上下文
	ctx := context.Background()
	// 基于root 创建一个 有退出状态的context
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	go task(cancelCtx)
	time.Sleep(time.Second * 3)
	fmt.Println("prepare cancelFunc ...")
	// 调用退出函数
	cancelFunc()
	fmt.Println("finish cancelFunc ...")
	time.Sleep(time.Second * 1)
	fmt.Println("end main goruntine ...")
}

func task(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully exit")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}
