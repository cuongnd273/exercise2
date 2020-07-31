package main

import (
	"context"
	"fmt"
	"time"
)

func work8() {
	t := time.NewTicker(100 * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-t.C:
				{
					fmt.Println("${time.Now().Unix()} done")
				}
			case <-ctx.Done():
				{
					done <- true
					return
				}
			}
		}
	}()
	<-done
	defer cancel()
}
