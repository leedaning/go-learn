package advance

import (
	"fmt"
	"sync"
)

func MySyncOnce2() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("onceBody")
	}

	done := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- i
		}()
	}
	for i := 0; i < 10; i++ {
		res := <-done
		fmt.Println("res", res)
	}
}
