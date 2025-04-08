package advance

import (
	"context"
	"fmt"
	"time"
)

// 通过context.WithCancel()创建一个可取消的context，调用返回的cancel函数可以取消该context
func MyContext3() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go worker3(ctx)
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func worker3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker: stopped")
			return
		default:
			fmt.Println("Worker: running")
			time.Sleep(1 * time.Second)
		}
	}
}
