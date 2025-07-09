package _select

import (
	"fmt"
	"sync"
	"time"
)

/*
	退出机制
	使用select实现优雅退出
	与context结合使用
*/

func SelectCase8() {
	//使用select实现服务退出
	quit := make(chan bool)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-quit:
				fmt.Println("Quit")
				return
			default:
				fmt.Println("Working......")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	quit <- true
	wg.Wait()
}
