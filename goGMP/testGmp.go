/*
 * @Descripttion:
 * @version:
 * @Author: xiaokun
 * @Date: 2022-04-13 20:20:51
 * @LastEditors: xiaokun
 * @LastEditTime: 2022-04-14 21:03:54
 */

package main

import (
	"fmt"
	"sync"

	"github.com/bigwhite/functrace"
	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

func main() {
	defer functrace.Trace()()
	log.Debugf("entry")
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go work(&wg)
	}
	wg.Wait()
}

func work(wg *sync.WaitGroup) {
	fmt.Println("worker!!")
	defer functrace.Trace()()
	cnt := 0
	Add(cnt)
	wg.Done()
}

func Add(n int) {
	defer functrace.Trace()()
	for i := 0; i < 1e10; i++ {
		n++
	}
}
