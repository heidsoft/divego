package main

import (
	"context"
	"fmt"
)

func main() {
	// 一级的context
	rootCtx := context.Background()
	fmt.Println(rootCtx)
	// 二级的context
	childCtx := context.WithValue(rootCtx, "msgId", "someMsgId")
	// 三级的context
	childOfChildCtx, cancelFunc := context.WithCancel(childCtx)
	fmt.Println(childOfChildCtx)
	fmt.Println(cancelFunc)
	// 一级比一级功能更全
}
