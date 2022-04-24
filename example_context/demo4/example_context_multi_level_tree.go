package main

import (
	"context"
	"fmt"
)

func main() {
	// 一级 context
	rootCtx := context.Background()
	// 二级 context，继承一级
	childCtx2 := context.WithValue(rootCtx, "msgId", "someMsgId")
	// 三级级 context，继承二级
	childCtx3, cancelFunc := context.WithCancel(childCtx2)
	fmt.Println(childCtx3)
	fmt.Println(cancelFunc)
	// 二级 context
	childCtx4 := context.WithValue(rootCtx, "user_id", "some_user_id")
	fmt.Println(childCtx4)
	fmt.Println(childCtx4.Value("user_id"))
}
