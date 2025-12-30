package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	chodd := make(chan bool)  // 用于同步信号
	cheven := make(chan bool) // 用于同步信号

	wg.Add(2)

	// 奇数goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= 99; i += 2 {
			<-cheven // 等待偶数信号
			fmt.Println("Odd:", i)
			chodd <- true // 发送信号给偶数
		}
	}()

	// 偶数goroutine
	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			<-chodd // 等待奇数信号
			fmt.Println("Even:", i)
			if i < 100 {
				cheven <- true // 发送信号给下一个奇数
			}
		}
	}()

	// 启动打印，从奇数开始
	cheven <- true
	wg.Wait()
	close(chodd)
	close(cheven)
	fmt.Println("Done printing 1 to 100 alternately.")
}
