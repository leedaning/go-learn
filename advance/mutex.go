package advance

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func MyMutex() {
	println("MyMutex, 互斥锁")
	wg := sync.WaitGroup{}
	var counter Counter

	/*for i := 0; i <= 100; i++ {
		wg.Add(1)
		go counter.Increment(&wg, i)
	}*/
	wg.Add(2)
	go mutexCircle(&wg, &counter)
	go mutexCircle(&wg, &counter)

	wg.Wait()
	println(counter.value)
}

func mutexCircle(wg *sync.WaitGroup, c *Counter) {
	println("开始调用循环计数方法")
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		//println(c.value)
		c.Increment(i)
	}
}

func (c *Counter) Increment(num int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	//c.value++
	c.value = c.value + num
}
