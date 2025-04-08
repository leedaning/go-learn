package advance

import (
	"fmt"
	"time"
)

/*
默认分支
default 分支是可选的，如果没有任何通道操作可以立即执行，default 分支会被执行。如果没有 default 分支，select 会一直阻塞，直到某个通道操作可以执行。
*/
func SelectCase3() {
	ch := make(chan int)

	/*wg := sync.WaitGroup{}
	wg.Add(1)*/
	go func() {
		//wg.Done()
		//time.Sleep(1 * time.Second)
		ch <- 1
	}()

	select {
	case v := <-ch:
		fmt.Println("Received value:", v)
	case <-time.After(3 * time.Second):
		fmt.Println("Timed out")
	default:
		fmt.Println("No value received")
	}

	// 继续执行其他代码
	time.Sleep(1 * time.Second)

	//wg.Wait()
}
