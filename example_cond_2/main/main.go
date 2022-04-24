package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	// 声明一把锁
	l := sync.Mutex{}
	c := sync.NewCond(&l)
	queue := make([]interface{},0,10)

	// 从队列移除
	removeFromQueue := func(delay time.Duration) {
		// 时间区间
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Lock()
		c.Signal()
	}

	for i:=0; i< 10; i++ {
		//条件锁
		c.L.Lock()
		for len(queue) == 2{
			//c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct {}{})
		//每隔一秒
		go removeFromQueue(1*time.Second)
		c.L.Unlock()
	}



}