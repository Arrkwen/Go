package main

import (
	"fmt"
	"time"
)

// channel 传递是传递channel指针的副本
// time.AfterFunc 是会开启新的子goroutine去执行func。
func delayFunc(i int, ch chan int) {
	time.AfterFunc(time.Second*time.Duration(2), func() {
		ch <- i
	})
}

// 普通基本类型传递是参数副本
func changeI(i int) {
	i -= 1
}

func server() {
	testch := make(chan int)

	for i := 0; i < 10; i++ {
		changeI(i)
		delayFunc(i, testch)
	}
	go func() {
		for {
			a, ok := <-testch
			if !ok {
				break
			}
			time.Sleep(time.Second * 1)
			fmt.Print(time.Now())
			fmt.Printf("I get the val %v\n", a)
		}
	}()
}

func TimeAfter() {
	go server()
	// 这儿是为了等待delayFUnc的子goroutine执行完毕，然后在server中的go func能消费到数据，不然主go routin直接退出，所有的子协程均退出
	for {
		time.Sleep(time.Second * 15)
		break
	}
}
