package main

import (
	"fmt"
	"sync"
)

const BATCHSIZE = 80
const GOROUTINENUM = 100

func process(k int, wg *sync.WaitGroup, foot [][]int) {
	SIZE := len(foot[0])
	fmt.Println("started Goroutine ", k)
	for m := 0; m < BATCHSIZE; m++ {
		for n := 0; n < SIZE; n++ {
			foot[m][n] = k
		}
	}
	fmt.Printf("Goroutine %d ended\n", k)
	wg.Done() // Done()用来表示goroutine已经完成了，减少一次计数器
}

func main() {

	size := 8000
	footTimeResult := make([][]int, size)
	for i := range footTimeResult {
		footTimeResult[i] = make([]int, size)
	}

	var wg sync.WaitGroup
	for i := 0; i < GOROUTINENUM; i++ {
		wg.Add(1)
		partitions := i * BATCHSIZE
		addres := footTimeResult[partitions : partitions+BATCHSIZE][:]
		go process(i, &wg, addres)
	}
	wg.Wait() // Wait()用来等待所有需要等待的goroutine完成。
	fmt.Println("All go routines finished executing")
	fmt.Println(footTimeResult)
}
