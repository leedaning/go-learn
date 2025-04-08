package advance

import (
	"fmt"
	"sync"
)

var counter = 0
var mutex sync.Mutex

func increment() {
	mutex.Lock()
	defer mutex.Unlock()
	counter++
	fmt.Println("counter:", counter)
}

func MyMutex2() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			//fmt.Println("goroutine 01:", i)
			increment()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			//fmt.Println("goroutine 02:", i)
			increment()
		}
	}()

	wg.Wait()
	fmt.Println("Final Counter2:", counter)
}
