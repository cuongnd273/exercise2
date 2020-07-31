package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func work3() {
	fmt.Println("*** Work 3 ***")
	t := time.NewTicker(3 * time.Second)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 9*time.Second)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			select {
			case <-t.C:
				{
					fmt.Println("Wait ")
				}
			case <-ctx.Done():
				{
					wg.Done()
					fmt.Println("Stop goroutine")
					return
				}
			}
		}
	}()
	wg.Wait()
	defer cancel()
}
