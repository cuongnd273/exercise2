package main

import (
	"fmt"
	"time"
)

func work1() {
	fmt.Println("*** Work 1 ***")
	ticker := time.NewTicker(3 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Printf("time now: {%s}\n", t)
			}
		}
	}()
	time.Sleep(9 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("kết thúc")
}
