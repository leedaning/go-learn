package advance

import (
	"fmt"
	"sync"
)

// 单例模式

var (
	once3    sync.Once
	instance *Singleton
)

type Singleton struct {
	val int
}

func GetInstance() *Singleton {
	once3.Do(func() {
		instance = &Singleton{
			val: 28,
		}
		fmt.Println("开启单例")
	})
	return instance
}

func MySyncOnce3() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			singleton := GetInstance()
			fmt.Printf("Goroutine %d: instance value: %d \n", id, singleton.val)
		}(i)
	}
	wg.Wait()
}
