package main

import (
	"fmt"
	"sync"
)

func main()  {
	// step 1 实例化cond, NewCond函数创建一个类型，满足sync,Locker接口。
	// 如此使得cond类型能够以一种并发安全的方式与其他goroutine协调
	c := sync.NewCond(&sync.Mutex{})
	c.L.Lock()
	for conditionTrue() == false{
		// 调用Wait不只阻塞，它挂起了当前的goroutine, 允许其他goroutine在OS线程上运行
		// 当调用Wait时，会发生一些其他事情
		// 进入Wait后，在Cond变量的Locker上调用Unlock方法
		fmt.Println("wait ...")
		c.Wait()//fatal error: all goroutines are asleep - deadlock!
	}
	c.L.Unlock()
}

func conditionTrue() bool {
	fmt.Println("条件判断...")
	return false
}