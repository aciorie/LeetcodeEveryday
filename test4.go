package main

import (
	"fmt"
	"time"
)

//当select监控多个chan同时到达就绪态时，
// 如何先执行某个任务？

// 可以在子case再加一个for select语句

func prioritySelect(ch1, ch2 <-chan string) {
	for {
		// 外层 select 优先检查 ch1
		select {
		case val := <-ch1:
			fmt.Printf("优先接收到：%s\n", val)
		default:
			// 如果 ch1 没有数据，则进入内层 select 等待 ch2 的数据
			select {
			case val2 := <-ch2:
				fmt.Printf("次优先接收到：%s\n", val2)
			case val := <-ch1: // 这里的 case 必须保留，以避免 ch1 产生死锁
				fmt.Printf("优先接收到：%s\n", val)
			}
		}
	}
}

func main4() {
	ch1 := make(chan string, 10)
	ch2 := make(chan string, 10)

	// 模拟并发发送数据
	go func() {
		for i := 0; i < 5; i++ {
			ch2 <- fmt.Sprintf("次级数据-%d", i)
			time.Sleep(100 * time.Millisecond) // 模拟 ch2 数据先到达
		}
	}()

	go func() {
		time.Sleep(250 * time.Millisecond) // 模拟 ch1 数据稍微延迟
		ch1 <- "高级数据-A"
		ch1 <- "高级数据-B"
	}()

	// 运行演示函数
	prioritySelect(ch1, ch2)
}
