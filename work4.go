package main

import (
	"fmt"
	"time"
)

func work4() {
	fmt.Println("*** Work 4 ***")
	t := time.Unix(0, 1592190294764144364)
	fmt.Println("Hour : ", t.Hour())
	fmt.Println("Minute : ", t.Minute())
	fmt.Println("Second : ", t.Second())
}
