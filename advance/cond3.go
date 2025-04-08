package advance

import (
	"fmt"
	"sync"
	"time"
)

func MyCond3() {

	//var shareData interface{}
	var ready bool
	ready = false
	fmt.Println(ready)

	cond3 := sync.NewCond(&sync.Mutex{})

	go func() {
		fmt.Println("等待方")
		// 等待方
		cond3.L.Lock()
		for !ready {
			fmt.Println("等待")
			cond3.Wait()
		}
		cond3.L.Unlock()
	}()

	time.Sleep(1 * time.Second)

	//通知方
	fmt.Println("通知方")
	cond3.L.Lock()
	ready = true
	cond3.Signal() // 或者cond.Broadcast()
	cond3.L.Unlock()
	time.Sleep(1 * time.Second)
	fmt.Println("End")
}
