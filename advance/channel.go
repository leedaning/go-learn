package advance

import (
	"fmt"
	"sync"
)

func MyChannel() {
	//channelBase()

	wg := sync.WaitGroup{}
	wg.Add(2)

	oddCh := make(chan int)
	evenCh := make(chan int)

	maxNum := 10
	go prtOddNumber(maxNum, oddCh, evenCh, &wg)
	go prtEvenNumber(maxNum, oddCh, evenCh, &wg)

	// 启动打印流程
	oddCh <- 0

	wg.Wait()

	close(oddCh)
	close(evenCh)
}

func channelBase() {
	ch := make(chan int)
	go func() {
		ch <- 18
	}()
	value := <-ch
	fmt.Println(value)
}

func prtOddNumber(end int, oddCh chan int, evenCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= end; i = i + 2 {
		<-oddCh // 等待信号
		fmt.Println(i)
		if i < end {
			evenCh <- 0 // 通知偶数goroutine
		}
	}
}

func prtEvenNumber(end int, oddCh chan int, evenCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= end; i = i + 2 {
		<-evenCh // 等待信号
		fmt.Println(i)
		if i < end {
			oddCh <- 0 // 通知奇数goroutine
		}
	}
}
