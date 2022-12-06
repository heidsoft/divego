package main

import (
	"fmt"
	"sync"
	"time"
)

// https://juejin.cn/post/7157986303085117470
func main() {
	c := make(chan string)

	//阻塞当前go 协程序，等待其他完成
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		//减少计数值
		defer wg.Done()
		c <- `foo`
		println(`发送消息成功`)
	}()

	go func() {
		defer wg.Done()

		time.Sleep(time.Second * 2)
		println(`收到消息: ` + <-c)
	}()

	//调用这个方法的 goroutine 会一直阻塞，直到 WaitGroup 的计数值变为 0
	//等待以上协程序完成
	fmt.Println("进入携程分组等待")
	//wg.Wait()
	//如果不加wg.Wait()等待，则程序退出，但是对应的go 子协程还在处理中
	fmt.Println("所有携程完成")
}
