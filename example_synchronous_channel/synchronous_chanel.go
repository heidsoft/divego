package main

import "fmt"

func main()  {
	ch := make(chan bool)

	go func() {
		ch <-true
	}()
	a :=<-ch

	fmt.Println(a)
}