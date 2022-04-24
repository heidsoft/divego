package main

import (
	"fmt"
	"time"
)

func main() {
	ch :=make(chan bool,1)

	ch<-true

	go func(chan bool) {
		fmt.Println("print ch")
		fmt.Println(<-ch)
	}(ch)

	time.Sleep(time.Second*5)
}
