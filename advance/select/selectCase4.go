package _select

import (
	"fmt"
)

// 非阻塞，我的理解就是有default，这样select就不会阻塞了
func SelectCase4() {
	ch := make(chan int)

	// 非阻塞发送
	select {
	case ch <- 28:
		fmt.Println("Value sent")
	default:
		fmt.Println("Channel is full")
	}
	/*wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wg.Done()
		ch <- 12
	}()
	wg.Wait()*/

	//非阻塞接收
	select {
	case v := <-ch:
		fmt.Println("Value received", v)
	default:
		fmt.Println("Channel is empty")
	}
}
