package advance

import (
	"fmt"
	"sync"
)

func Coroutine() {

	var wg sync.WaitGroup
	startRoutine(&wg, 1, 5)
	startRoutine(&wg, 1, 10)
	startRoutine(&wg, 10, 20)
	wg.Wait()
}

func startRoutine(wg *sync.WaitGroup, begin, end int) {

	wg.Add(1)
	go circle(wg, begin, end)
}

func circle(wg *sync.WaitGroup, begin, end int) {
	defer wg.Done()
	if begin >= end {
		return
	}
	temp := 0
	for i := begin; i <= end; i++ {
		temp += i
	}
	fmt.Println(temp)
	//return temp
}
