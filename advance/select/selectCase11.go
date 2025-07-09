package _select

import (
	"fmt"
)

// select {} // 永久阻塞，常用于main函数防止程序退出
func SelectCase11() {
	fmt.Println("selectCase11")
	select {
	//case <-time.After(10 * time.Second):
	//	fmt.Println("timeout")
	} // 永久阻塞，常用于main函数防止程序退出
	fmt.Println("selectCase11 end")
}
