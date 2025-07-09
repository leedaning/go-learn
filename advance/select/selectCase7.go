package _select

import "fmt"

func SelectCase7() {
	ch := make(chan int, 1)

	for i := 0; i < 10; i++ {
		select {
		case res := <-ch:
			fmt.Println(res)
		//case ch <- i:
		default:
			fmt.Println("Default:", i)
			ch <- i
		}
	}
}
