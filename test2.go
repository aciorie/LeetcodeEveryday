package main

import (
	"fmt"
	"time"
)

// 两个协程交替打印10个字母和数字
var word = make(chan struct{}, 1)
var num = make(chan struct{}, 1)

func printLetter() {
	for i := 0; i < 10; i++ {
		<-word
		fmt.Printf("%c\n", 'a'+i)
		num <- struct{}{}
	}
}

func printNum() {
	for i := 0; i < 10; i++ {
		<-num
		fmt.Println(i)
		word <- struct{}{}
	}
}

func main2() {
	go printLetter()
	go printNum()
	word <- struct{}{}
	time.Sleep(2 * time.Second)
}
