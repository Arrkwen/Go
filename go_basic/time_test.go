package main

import (
	"fmt"
	"time"
)

func TimeDuration() {
	t1 := time.Now()
	time.Sleep(time.Second * 7)
	if time.Now().Sub(t1) < time.Second*time.Duration(6) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
