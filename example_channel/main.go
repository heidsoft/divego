package main

import (
	"fmt"
	"time"
)

func goRoutineA(a <-chan int)  {
	val := <-a
	fmt.Println("goRoutineA received the data", val)
}
func main()  {
	ch := make(chan int)
	go goRoutineA(ch)
	ch <-1
	time.Sleep(time.Second * 1)
}


