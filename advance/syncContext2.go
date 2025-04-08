package advance

import (
	"context"
	"fmt"
	"time"
)

// 使用context.Background()创建一个空的context，作为其他context的根
func MyContext2() {
	ctx := context.Background()
	go worker2(ctx)

	time.Sleep(time.Second * 2)
}

func worker2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work canceled.")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}

	}
}
