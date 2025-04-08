package advance

import (
	"fmt"
	"sync"
)

func MyWaitGroup() {

	wg := new(sync.WaitGroup)
	wg.Add(2)
	ch := make(chan int)
	ch2 := make(chan int)
	defer close(ch)
	go countNum(1, 10, ch, wg)
	go countNum(2, 10, ch2, wg)

	res := <-ch
	res2 := <-ch2
	wg.Wait()
	println("奇数和：", res)
	println("偶数和：", res2)
	
}

func countNum(start int, end int, ch chan int, wg *sync.WaitGroup) int {
	defer wg.Done()
	temp := 0
	for i := start; i <= end; i = i + 2 {
		temp += i
	}
	ch <- temp
	fmt.Println("总数：", temp)
	return temp
}
