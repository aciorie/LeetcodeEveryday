package main

import (
	"fmt"
	"sync"
)

/*
3两个协程交替打印10个字母和数字
思路：采用channel来协调goroutine之间顺序。

主线程一般要waitGroup等待协程退出，这里简化了一下直接sleep。
*/
func main4() {
	chanNum, chanChar := make(chan bool), make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			<-chanNum
			fmt.Println("%d", i)
			chanChar <- true
		}
	}()

	go func() {
		defer wg.Done()
		for i := 'A'; i <= 'J'; i++ {
			<-chanChar
			fmt.Println("%c", i)
			chanNum <- true
		}
	}()

	chanNum <- true
	wg.Wait()
	fmt.Println("Done")
}
