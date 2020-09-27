package main

import (
	"fmt"
)
// 共享内存测试
func addByShareCommunicate(n int) []int{
	var ints []int
	channel :=make(chan int,n)

	//写chan
	for i :=0; i < n; i++{
		go func(channel chan<- int, order int) {
			//写入管道
			channel <-order
		}(channel,i)
	}

	for i := range channel{
		//读取管道
		ints = append(ints,i)
		if len(ints) == n{
			break
		}
	}
	close(channel)
	return ints
}

func main()  {
	foo := addByShareCommunicate(10)
	fmt.Println(len(foo))
	fmt.Println(foo)
}