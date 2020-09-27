package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
func main()  {
	wg.Add(1)
	fmt.Println("start ")
	go func() {
		fmt.Println("go func")
		handle1()
		wg.Done()
	}()

	//等待以上结束
	wg.Add(1)
	go handle2()
	wg.Wait()
}

func handle1()  {
	for i:=0; i<4; i++{
		fmt.Printf("handle1 %d\n", i)
		time.Sleep(time.Second)
	}
}

func handle2()  {
	for i:=0; i<4; i++{
		fmt.Printf("handle2 %d\n", i)
		time.Sleep(time.Second)
	}
	wg.Done()
}
