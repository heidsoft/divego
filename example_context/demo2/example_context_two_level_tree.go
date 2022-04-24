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
	fmt.Println(childCtx)
	fmt.Println(childCtx.Value("msgId"))
}
