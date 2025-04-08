package advance

import (
	"fmt"
	"sync"
	"time"
)

/*
随机选择
如果多个 case 中的通道操作都可以执行，select 会随机选择其中一个执行。这种随机性确保了公平性。
*/
func SelectCase() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		wg.Done()
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	go func() {
		wg.Done()
		time.Sleep(1 * time.Second)
		ch2 <- 2
	}()

	//for i := 0; i < 2; i++ {

	select {
	case v := <-ch1:
		fmt.Println("Received from ch1: ", v)
	case v := <-ch2:
		fmt.Println("Received from ch2: ", v)
	}
	//time.Sleep(1 * time.Second)
	//}
	wg.Wait()
}
