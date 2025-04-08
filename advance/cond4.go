package advance

import (
	"fmt"
	"sync"
	"time"
)

/*生产者-消费者*/
type queue struct {
	data []int      // 存储数据的切片
	cap  int        // 队列容量
	cond *sync.Cond // 条件变量，用于协调生产者和消费者
}

func newQueue4(cap int) *queue {
	q := &queue{
		cap:  cap,
		cond: sync.NewCond(&sync.Mutex{}), // 使用互斥锁创建条件变量
	}
	return q
}

func (q *queue) push(item int) error {
	fmt.Println("生产者")
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	for len(q.data) == q.cap {
		fmt.Println("生产者-开始等待")
		q.cond.Wait() // 队列满时等待
	}

	q.data = append(q.data, item)
	fmt.Printf("生产者-添加 %d, 队列: %v\n", item, q.data)
	q.cond.Broadcast() // 唤醒可能等待的消费者
	//q.cond.Signal() // 唤醒一个等待的消费者
	return nil
}

func (q *queue) pop() (int, error) {
	fmt.Println("\n消费者")
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	for len(q.data) == 0 {
		fmt.Println("消费者-队列为空，等待")
		q.cond.Wait()
	}

	item := q.data[0]
	q.data = q.data[1:]
	fmt.Printf("消费者-取出 %d, 队列: %v\n", item, q.data)
	q.cond.Broadcast() // 唤醒可能等待的生产者
	//q.cond.Signal() // 唤醒一个等待的生产者
	return item, nil
}

func MyCond4() {
	q := newQueue4(2)

	func() {
		q.cond.L.Lock()
		defer q.cond.L.Unlock()
		q.data = append(q.data, 1)
		q.data = append(q.data, 2)
	}()

	var wg sync.WaitGroup

	//生产者
	wg.Add(1)
	go func() {
		defer wg.Done()
		q.push(3)
		q.push(4)
		for i := 5; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			q.push(i)
		}
	}()

	for i := 0; i < 15; i++ {
		// 消费者
		wg.Add(1)
		go func() {
			defer wg.Done()
			item, err := q.pop()
			if err != nil {
				fmt.Println("消费错误：", err)
				return
			}
			fmt.Println("成功消费：", item)
		}()
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
}
