package advance

import (
	"context"
	"fmt"
	"time"
)

func MyContext() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	go worker(ctx)

	time.Sleep(time.Second * 5)
	fmt.Println("Main goroutine finished")
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return
		default:
			fmt.Println("Worker running")
			time.Sleep(1 * time.Second)
		}

	}
}
