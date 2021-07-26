package main

//非阻塞的收发
//在通常情况下，select 语句会阻塞当前 Goroutine 并等待多个 Channel 中的一个达到可以收发的状态。
//但是如果 select 控制结构中包含 default 语句，那么这个 select 语句在执行时会遇到以下两种情况：
//当存在可以收发的 Channel 时，直接处理该 Channel 对应的 case；
//当不存在可以收发的 Channel 时，执行 default 中的语句；
//当我们运行下面的代码时就不会阻塞当前的 Goroutine，它会直接执行 default 中的代码。
//https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-select/
func main() {
	ch := make(chan int)
	select {
		case i := <-ch:
			println(i)

		//default:
		//	println("default")
	}
}

