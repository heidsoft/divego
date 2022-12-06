package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// https://stackoverflow.com/questions/68593779/can-unbuffered-channel-be-used-to-receive-signal
// https://github.com/golang/go/issues/45604
// https://lailin.xyz/post/go-training-week3-waitgroup.html
// https://www.mohitkhare.com/blog/waitgroups-in-golang/
// https://www.henrydu.com/2020/10/27/golang-channels/
func main() {
	//无缓冲管道
	signals := make(chan os.Signal, 1)
	//signals := make(chan os.Signal)
	//SIGINT 程序终止(interrupt)信号, 在用户键入INTR字符(通常是Ctrl-C)时发出，用于通知前台进程组终止进程
	//SIGTERM 程序结束(terminate)信号, 与SIGKILL不同的是该信号可以被阻塞和处理。通常用来要求程序自己正常退出，shell命令kill缺省产生这个信号。如果进程终止不了，我们才会尝试SIGKILL。
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case sig := <-signals:
			fmt.Println("信号名称", sig.String())
		default:
			fmt.Println("当前时间", time.Now())
		}
		time.Sleep(2 * time.Second)
	}
}
