package advance

import (
	"context"
	"fmt"
	"time"
)

/*
结合 context 使用
select 常用于结合 context 来处理超时或取消操作。
*/
func SelectCase5() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 28
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Context done:", ctx.Err())
	case v := <-ch:
		fmt.Println("Received value:", v)
	}
}
