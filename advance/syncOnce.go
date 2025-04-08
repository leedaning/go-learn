package advance

import (
	"fmt"
	"sync"
)

var count int
var once sync.Once

func MySyncOnce() {
	once.Do(incrementOnce)
	for i := 0; i < 10; i++ {

		once.Do(incrementOnce)
		incrementOnce()
	}
	once.Do(incrementOnce)
	fmt.Println(count)
}

func incrementOnce() {
	count++
}
