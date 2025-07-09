package _select

import (
	"fmt"
	"time"
)

/*
优先级控制
可以通过select的嵌套模式实现case优先级控制
*/
func SelectCase10() {
	highPriority := make(chan string)
	lowPriority := make(chan string)

	//wg := sync.WaitGroup{}
	//wg.Add(2)
	go func() {
		//wg.Done()
		lowPriority <- "low"
	}()

	go func() {
		//wg.Done()
		time.Sleep(1 * time.Second)
		highPriority <- "high"
	}()
	//wg.Wait()
	//time.Sleep(1001 * time.Millisecond)

	select {
	case msg := <-highPriority:
		fmt.Println("High priority value:", msg)
	default:
		time.Sleep(1000 * time.Millisecond)
		select {
		case msg := <-highPriority:
			fmt.Println("High priority value2:", msg)
		case msg := <-lowPriority:
			fmt.Println("Low priority value:", msg)
		default:
			fmt.Println("No Priority")
		}
	}
}
