package advance

import (
	"fmt"
)

func Coroutine2() {
	startRoutine2()
}

func startRoutine2() {
	for i := 0; i <= 10; i = i + 5 {
		resChan := make(chan int)
		defer close(resChan)
		go circle2(resChan, 1, i)
		res := <-resChan
		fmt.Println(res)
	}

}

func circle2(resChan chan int, begin, end int) {
	if begin >= end {
		resChan <- 0
		return
	}
	temp := 0
	for i := begin; i <= end; i++ {
		temp += i
	}
	resChan <- temp
}
