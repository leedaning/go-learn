package basic

import (
	"bufio"
	"fmt"
	"os"
)

/*
统计标准输入中每行文本的出现次数，并输出那些重复出现的行及其出现次数
*/
func SearchRepeat() {
	counts := make(map[string]int)      // 创建一个 map，用于存储每行文本的出现次数
	input := bufio.NewScanner(os.Stdin) // 创建一个 Scanner，从标准输入读取数据
	for input.Scan() {                  // 循环读取每一行
		counts[input.Text()]++ // 将当前行的计数加1
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts { // 遍历 map
		if err := input.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
		}
		if n > 1 { // 如果某行的出现次数大于1
			fmt.Printf("%s:\t%d\n", line, n) // 输出该行及其出现次数
		}
	}
}
