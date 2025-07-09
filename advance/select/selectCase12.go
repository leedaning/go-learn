package _select

import (
	"fmt"
	"reflect"
)

// 反射与select
func SelectCase12() {

	//使用反射实现动态select
	ch1 := make(chan int, 1)
	ch2 := make(chan string, 1)

	ch1 <- 18
	ch2 <- "hello"

	cases := []reflect.SelectCase{
		{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch1)},
		{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch2)},
	}

	chosen, value, _ := reflect.Select(cases)
	fmt.Printf("CHosen %d, value: %v\n", chosen, value)
}
