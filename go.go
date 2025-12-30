package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan bool) // 用于同步信号

	wg.Add(2)

	// 奇数goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= 99; i += 2 {
			fmt.Println("Odd:", i)
			ch <- true // 发送信号给偶数
		}
	}()

	// 偶数goroutine
	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			<-ch // 等待奇数信号
			fmt.Println("Even:", i)
			if i < 100 {
				ch <- true // 发送信号给下一个奇数
			}
		}
	}()

	// 开始：奇数先打印，无需初始信号
	wg.Wait()
	close(ch)
	fmt.Println("Done printing 1 to 100 alternately.")
}
