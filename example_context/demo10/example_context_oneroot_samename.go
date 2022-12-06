package main

import (
	"context"
	"fmt"
)

type keyOne string
type keyTwo string

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, keyOne("one"), "valueOne")
	ctx = context.WithValue(ctx, keyTwo("one"), "valueTwo")

	fmt.Println(ctx.Value(keyOne("one")).(string))
	fmt.Println(ctx.Value(keyTwo("one")).(string))
}
