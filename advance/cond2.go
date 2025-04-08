package advance

import (
	"fmt"
	"sync"
	"time"
)

/*
生产者-消费者
*/

type Queue struct {
	cond         *sync.Cond
	queue        []interface{}
	shuttingDown bool
}

func newQueue() *Queue {
	return &Queue{
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

func (q *Queue) add(item interface{}) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if q.shuttingDown {
		return
	}
	q.queue = append(q.queue, item)
	q.cond.Signal()
}

func (q *Queue) get() (interface{}, bool) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for len(q.queue) == 0 && !q.shuttingDown {
		q.cond.Wait()
	}
	if len(q.queue) == 0 {
		return nil, true
	}
	item := q.queue[0]
	q.queue = q.queue[1:]
	return item, false
}

func (q *Queue) ShutDown() {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.shuttingDown = true
	q.cond.Broadcast()
}

func MyCond2() {
	queue := newQueue()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				item, shutdown := queue.get()
				if shutdown {
					break
				}
				fmt.Printf("Goroutine %d: received item %v\n", id, item)
			}
		}(i)
	}

	go func() {
		for i := 10; i >= 1; i-- {
			queue.add(i)
			time.Sleep(100 * time.Millisecond)
		}
		queue.ShutDown()
	}()

	wg.Wait()
}
