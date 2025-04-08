package advance

import (
	"fmt"
	"sync"
)

func MyChannel2() {

	wg := sync.WaitGroup{}
	wg.Add(2)

	oddCh := make(chan int)
	evenCh := make(chan int)

	maxNum := 10
	go prtOddNum(maxNum, oddCh, evenCh, &wg)
	go prtEvenNum(maxNum, oddCh, evenCh, &wg)

	oddCh <- 0

	wg.Wait()

	close(oddCh)
	close(evenCh)

}

// 打印奇数
func prtOddNum(maxNum int, oddCh chan int, evenCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= maxNum; i = i + 2 {
		<-oddCh
		fmt.Println(i)
		if i+1 <= maxNum {
			evenCh <- 0
		}
	}
}

// 打印偶数
func prtEvenNum(maxNum int, oddCh chan int, evenCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= maxNum; i = i + 2 {
		<-evenCh
		fmt.Println(i)
		if i+1 <= maxNum {
			oddCh <- 0
		}
	}

}
