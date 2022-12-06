package main

import (
	"context"
	"fmt"
)

//ctxtmp := context.WithValue(context.Background(), "1111", "2222")
//第一种可取值：tidandsid1, _ := ctxtmp.Value("1111").(string)
//第一种无法取值：tidandsid2, _ := context.Background().Value("1111").(string)
//第二种怎么才能获取到值呢？

type ctxKey string
type ctxVal string

func main() {
	// 构建根root
	rootCtx := context.Background()
	// 构建能存储值的ctx
	rootCtx = context.WithValue(rootCtx, "1111", "2222")
	// 获取 kvCtx的 key value
	//fmt.Println(kvCtx)

	//
	fmt.Println(rootCtx.Value(ctxKey("1111")))

}
