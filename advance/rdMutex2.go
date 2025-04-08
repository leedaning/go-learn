package advance

import (
	"fmt"
	"sync"
	"time"
)

/*
未使用读写锁的示例
*/
type Counter2 struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter2 {
	return &Counter2{}
}

// 读计数器
func (c *Counter2) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// 更新计数器
func (c *Counter2) UpdateValue(delta int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += delta
}

func MyRDMutex2() {
	counter := NewCounter()
	var wg sync.WaitGroup

	//启动多个读取goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				start := time.Now().UnixNano()
				value := counter.GetValue()
				end := time.Now().UnixNano()
				//fmt.Printf("Reader %d: value = %d, time = %s\n", id, value, time.Now().Format("2006-01-02 15:04:05.000"))
				fmt.Printf("Reader %d: value = %d, time = %dns\n", id, value, end-start)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	// 启动一个更新goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			start := time.Now().UnixNano()
			counter.UpdateValue(1)
			end := time.Now().UnixNano()
			//fmt.Printf("Updater: updated value to %d, time = %s\n", counter.GetValue(), time.Now().Format("2006-01-02 15:04:05.000"))
			fmt.Printf("Updater: updated value to %d, time = %dns\n", counter.GetValue(), end-start)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	wg.Wait()
}
