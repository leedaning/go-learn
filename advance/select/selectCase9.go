package _select

import (
	"fmt"
	"time"
)

/*
定时任务
使用Ticker实现定时任务
定时器与select结合
*/
func SelectCase9() {
	//使用Ticker实现定时任务
	ticker := time.NewTicker(1000 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}

		}
	}()

	time.Sleep(10 * time.Second)
	ticker.Stop()
	done <- true
}
