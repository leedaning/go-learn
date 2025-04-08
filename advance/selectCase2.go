package advance

import (
	"fmt"
	"time"
)

/*
超时处理
select 可以结合 time.After 来实现超时处理。如果在指定时间内没有通道操作完成，default 分支会被执行。
*/
func SelectCase2() {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()

	select {
	case v := <-ch:
		fmt.Println("Received value:", v)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout 2 second")
	}
}
