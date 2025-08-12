package main

import (
	"fmt"
	"time"
)

/*
2（Goroutine）有三个函数，分别打印"cat", "fish","dog"
要求每一个函数都用一个goroutine，按照顺序打印100次。
此题目考察channel，用三个无缓冲channel，如果一个channel收到信号
则通知下一个。
*/

var cat = make(chan struct{})
var fish = make(chan struct{})
var dog = make(chan struct{})

func Cat() {
	<-cat
	fmt.Println("cat")
	fish <- struct{}{}
}

func Fish() {
	<-fish
	fmt.Println("fish")
	dog <- struct{}{}
}

func Dog() {
	<-dog
	fmt.Println("dog")
	cat <- struct{}{}
}

func main3() {
	for i := 0; i < 100; i++ {
		go Cat()
		go Fish()
		go Dog()
	}
	cat <- struct{}{}
	time.Sleep(2 * time.Second)
}
