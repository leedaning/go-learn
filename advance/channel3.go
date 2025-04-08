package advance

import (
	"fmt"
	"sync"
)

func MyChannel3() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	oddCh := make(chan int)
	evenCh := make(chan int)

	max := 10
	go ji(max, oddCh, evenCh, &wg)
	go ou(max, oddCh, evenCh, &wg)

	// 启动打印流程
	oddCh <- 0

	wg.Wait()

	close(oddCh)
	close(evenCh)
}

func ji(max int, oddCh chan int, evenCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= max; i = i + 2 {
		<-oddCh // 等待信号
		fmt.Printf("%d\n", i)
		if i < max {
			evenCh <- 0 // 通知偶数 goroutine
		}
	}
}

func ou(max int, oddCh chan int, evenCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= max; i = i + 2 {
		<-evenCh // 等待信号
		fmt.Printf("%d\n", i)
		if i < max {
			oddCh <- 0 // 通知奇数 goroutine
		}
	}
}
