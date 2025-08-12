package main

import "fmt"

/*
5当select监控多个chan同时到达就绪态时，如何先执行某个任务？
可以在子case再加一个for select语句。
*/
func priority_select(ch1, ch2 <-chan string) {
	for {
		select {
		case val := <-ch1:
			fmt.Println(val)
		case val2 := <-ch2:
		priority:
			for {
				select {
				case val1 := <-ch1:
					fmt.Println(val1)

				default:
					break priority
				}
			}
			fmt.Println(val2)
		}
	}
}
