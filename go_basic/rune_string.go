package main

import (
	"fmt"
	"unsafe"
)

// type Stu struct {
// 	Name string `json:"stu_name"`
// 	ID   string `json:"stu_id"`
// 	Age  int    `json:"-"`
// }

// func main1() {
// 	buf, _ := json.Marshal(Stu{"Tom", "t001", 18})
// 	fmt.Printf("%s\n", buf)
// }

func StringAndRune() {
	str := "go语言"
	fmt.Println(len(str))           //len是返回字节个数
	fmt.Println(unsafe.Sizeof(str)) // sizeof 是返回内存占用，而字符串的内存包括：指针+长度+实际的字符数
	fmt.Println("------")
	runeStr := []rune(str)
	fmt.Println(len(runeStr)) // 返回字符数
	fmt.Println(unsafe.Sizeof(runeStr))
	fmt.Println("------")
	for i, val := range str {
		fmt.Println(str[i])     //输出前四字节对应的unicode数字
		fmt.Println(val)        //输出unicode数字
		fmt.Printf("%c\n", val) //输出字符
	}
	fmt.Println("------")
	for i, val := range runeStr {
		fmt.Println(str[i])
		fmt.Println(val)
		fmt.Printf("%c\n", val)
	}
	s := "GÖ"
	sample := "H哈"

	sByte := []byte(s)
	sRune := []rune(s)
	sampleByte := []byte(sample)
	sampleRune := []rune(sample)

	fmt.Printf("%s\nsByte: %d\nsRune: %d\n", s, sByte, sRune)
	fmt.Println("------")
	fmt.Printf("%s\nsampleByte: %d\nsampleRune: %d\n", sample, sampleByte, sampleRune)
}
