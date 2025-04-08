package advance

import (
	"fmt"
	"sync"
	"time"
)

/*
使用读写锁
*/
type Counter3 struct {
	mu    sync.RWMutex
	value int
}

// 新建计数器
func NewCounter3() *Counter3 {
	return &Counter3{}
}

// 读计数器
func (c *Counter3) GetValue() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

// 更新计数器
func (c *Counter3) SetValue(delta int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += delta
}

func MyRDMutex3() {
	counter := NewCounter3()
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
				//fmt.Printf("Reader %d: value:%d, time=%s \n", id, value, time.Now().Format("2006-01-02 15:04:05.000"))
				fmt.Printf("Reader %d: value:%d, time=%dns \n", id, value, end-start)
				time.Sleep(100 * time.Millisecond)
			}
			counter.GetValue()
		}(i)
	}

	// 启动一个更新goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {

			start := time.Now().UnixNano()
			counter.SetValue(1)
			end := time.Now().UnixNano()
			//fmt.Printf("Updater: updated value to %d, time = %d\n", counter.GetValue(), time.Now().Format("2006-01-02 15:04:05.000"))
			fmt.Printf("Updater: updated value to %d, time = %dns\n", counter.GetValue(), end-start)
			time.Sleep(500 * time.Millisecond)
		}
	}()
	wg.Wait()
}
