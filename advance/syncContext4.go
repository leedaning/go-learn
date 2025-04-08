package advance

import (
	"context"
	"fmt"
	"time"
)

//通过context.WithTimeout()创建一个在指定时间后自动取消的context

func MyContext4() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go worker4(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("Worker completed or timeout")
	}
}

func worker4(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker done")
			return
		default:
			fmt.Println("worker running")
			return
		}
	}
}
