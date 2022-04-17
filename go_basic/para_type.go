package main

import (
	"fmt"
	"reflect"
)

func ParaType() {
	// map  channel slice func
	m := make(map[string]string)
	m["脑子进煎鱼了"] = "这次一定！"
	fmt.Printf("the map type:%v\n", reflect.TypeOf(m))
	fmt.Printf("the map content:%v\n", reflect.ValueOf(m))
	fmt.Printf("main:the value of m：%p\n", m)
	fmt.Printf("main: the addres of m：%p\n", &m)
	mapAfterFunc(m)
	fmt.Printf("%v\n", m)

	// array
	a := [3]int{1, 2, 3}
	fmt.Printf("the  array type:%v\n", reflect.TypeOf(a))
	fmt.Printf("the array content:%v\n", reflect.ValueOf(a))
	fmt.Printf("the value of a：%v\n", a)
	fmt.Printf("the address of a：%p\n", &a)
	arrayAfterFunc(a)
	fmt.Printf("%v\n", a)

	// string
	s := "abcde"
	fmt.Printf("the type of s:%v\n", reflect.TypeOf(s))
	fmt.Printf("the content of s:%v\n", reflect.ValueOf(s))
	fmt.Printf("the value of s：%s\n", s)
	fmt.Printf("the address of s：%p\n", &s)
	stringAfterFunc(s)
	fmt.Printf("%v\n", s)
}

func mapAfterFunc(m map[string]string) {
	fmt.Printf("hello: the value of m：%p\n", m)
	fmt.Printf("hello:the address of m：%p\n", &m)
	m["脑子进煎鱼了"] = "记得点赞！"
}

func arrayAfterFunc(a [3]int) {
	a[0] = 4
}

func stringAfterFunc(s string) {
	for _, val := range s {
		val += 1
	}
}
