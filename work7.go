package main

import (
	"context"
	"fmt"
	"time"
)

func x(ctx context.Context) {
	val := ctx.Value("time")
	if val == nil {
		fmt.Println("Not found")
	} else {
		t := time.Now().UnixNano()
		result := t - val.(int64)
		fmt.Println("Result :", result)
	}
}
func work7() {
	t := time.Now().UnixNano()
	ctx := context.WithValue(context.Background(), "time", t)
	x(ctx)
}
