package advance

import (
	"fmt"
	"sync"
)

/*
	读写互斥锁
*/

type SafeData struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewSafeData() *SafeData {
	return &SafeData{
		data: make(map[string]int),
	}
}

// 读操作
func (sd *SafeData) Read(key string) int {
	sd.mu.RLock()
	defer sd.mu.RUnlock()
	return sd.data[key]
}

func (sd *SafeData) Write(key string, value int) {
	sd.mu.Lock()
	defer sd.mu.Unlock()
	sd.data[key] = value
}

func MyRDMutex() {
	sd := NewSafeData()

	//启动多个goroutine进行读写操作
	var wg sync.WaitGroup

	// 写操作
	wg.Add(1)
	go func() {
		defer wg.Done()
		sd.Write("counter", 28)
	}()
	//wg.Wait()

	// 读操作
	wg.Add(1)
	go func() {
		defer wg.Done()
		val := sd.Read("counter")
		fmt.Println("读出的数据：", val)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		val := sd.Read("counter")
		fmt.Println("读出的数据为：", val)
	}()

	wg.Wait()
	fmt.Println(sd.data)
}
