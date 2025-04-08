package advance

import (
	"fmt"
	"sync"
	"time"
)

//条件变量Cond

func MyCond() {
	c := sync.NewCond(&sync.Mutex{})
	condition := false

	go func() {
		fmt.Println("Wait before condition")
		c.L.Lock()
		for !condition {
			c.Wait()
		}
		fmt.Println("Condition met, continue execution")
		c.L.Unlock()
		fmt.Println("Wait after condition")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Signal before")
	c.L.Lock()
	condition = true
	c.L.Unlock()
	c.Signal()
	fmt.Println("Signal after")
	time.Sleep(1 * time.Second)
}
