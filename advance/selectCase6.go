package advance

import (
	"fmt"
	"time"
)

/*
多路复用
select 可以用于多路复用，监听多个通道的操作，确保公平性。
select 语句会随机选择一个就绪的通道操作来执行，如果多个通道操作都就绪，则会随机选择其中一个。
*/
func SelectCase6() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 1; i < 10; i++ {
			ch1 <- i
			time.Sleep(time.Millisecond * 100)
		}
		close(ch1)
	}()

	go func() {
		for i := 1; i < 10; i++ {
			ch2 <- i * 10
			time.Sleep(time.Millisecond * 100)
		}
		close(ch2)
	}()

	for {

		select {
		case v, ok := <-ch1: //  ok 用于判断通道是否已经关闭。如果 ok 为 false，说明通道已经关闭。
			if !ok {
				fmt.Println("ch1 closed")
				return
			}
			fmt.Println("Received from ch1 value:", v)
		case v, ok := <-ch2:
			if !ok {
				fmt.Println("ch2 closed")
				return
			}
			fmt.Println("Received from ch2 value:", v)
		}
	}
}
