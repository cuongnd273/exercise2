package main

import (
	"fmt"
	"time"
)

func work5() {
	fmt.Println("*** Work 5 ***")
	t := time.Unix(1592190294764144364, 0)
	fmt.Println("Day of week : ", t.Weekday())
}
