package main

import (
	"fmt"
	"sync"
)

func main2() {
	var count int         // 普通的 int 变量
	var mu sync.Mutex     // 保护 count 变量的互斥锁
	var wg sync.WaitGroup // WaitGroup 用于等待所有 Goroutine 完成

	numWorkers := 3              // 启动 3 个 Goroutine
	incrementsPerWorker := 10000 // 每个 Goroutine 增加的次数

	fmt.Println("--- 不自定义计数器类型，直接使用 sync.Mutex 保护 int 变量 ---")
	fmt.Printf("启动 %d 个 Goroutine，每个 Goroutine 增加计数 %d 次\n", numWorkers, incrementsPerWorker)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d: 开始计数...\n", workerID)
			for j := 0; j < incrementsPerWorker; j++ {
				mu.Lock() // 加锁
				count++   // 增加 count
				fmt.Println("id: %d , count: %d ", workerID, count)
				mu.Unlock() // 解锁
				// time.Sleep(time.Microsecond) // 模拟一些工作，让并发冲突更明显
			}
			fmt.Printf("Goroutine %d: 计数完成。\n", workerID)
		}(i + 1) // 从 1 开始编号 Goroutine
	}

	wg.Wait() // 等待所有 Goroutine 完成

	// 读取最终计数值时也需要加锁，以确保读取到的是最新、最准确的值
	mu.Lock()
	finalCount := count
	mu.Unlock()

	fmt.Printf("\n最终计数: %d\n", finalCount)
	fmt.Printf("预期计数: %d\n", numWorkers*incrementsPerWorker)
	fmt.Println("--------------------------------------")
}
