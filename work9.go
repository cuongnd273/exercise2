package main

import (
	"fmt"
	"time"
)

func work9() {
	done := make(chan bool)
	time.AfterFunc(1*time.Second, func() {
		fmt.Println("i'm study")
		done <- true
	})
	<-done
}
