package basic

import (
	"fmt"
	"os"
	"strings"
)

func MyFor() {
	// 方法一、通过循环，+=的方式连接，如果数据量很大，代价高昂
	/*s, sep := "", ""                  // 初始化两个变量，s 用于存储最终结果，sep 用于分隔符
	for _, arg := range os.Args[1:] { // 遍历 os.Args[1:]，即从第二个参数开始的所有参数
		s += sep + arg // 将当前参数 arg 拼接到 s 中，sep 是分隔符
		sep = " "      // 第一次循环后，sep 被设置为空格，用于后续参数的分隔
	}
	fmt.Println(s) // 输出拼接后的字符串*/

	// 方法二、使用strings包的Join函数
	fmt.Println(strings.Join(os.Args[0:1], " ")) // 这种更简单高效
}
