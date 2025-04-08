package main

import "MyGo/advance"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	//s := "gopher"
	//fmt.Println("Hello and welcome, %s!", s)

	/*for i := 1; i <= 5; i++ {
		//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		fmt.Println("i =", 100/i)
	  }*/

	//advance.Coroutine()
	//advance.Coroutine2()

	//advance.MyChannel()
	//advance.MyChannel2()
	//advance.MyChannel3()
	//advance.MyChannel4()

	//advance.MyMutex()
	//advance.MyMutex2()

	//advance.MyRDMutex()	// 读写锁
	//advance.MyRDMutex2() // 未使用读写锁
	//advance.MyRDMutex3() // 使用读写锁

	//advance.MyWaitGroup()		// 等待组

	//advance.MySyncOnce() // 一次执行
	//advance.MySyncOnce2() //
	//advance.MySyncOnce3() // 单例模式
	//advance.MySyncOnce4() // 单例模式
	//advance.MySyncOnce5() // 延迟初始化
	//advance.MySyncOnce6() // 处理带错误的once

	//advance.MyCond() // 条件变量（Cond）
	//advance.MyCond2() //生产者-消费者
	//advance.MyCond3()
	//advance.MyCond4()

	//advance.MySyncMap()

	//advance.MyContext()
	//advance.MyContext2() //使用context.Background()创建一个空的context
	//advance.MyContext3() //通过context.WithCancel()创建一个可取消的context，调用返回的cancel函数可以取消该context
	//advance.MyContext4() //通过context.WithTimeout()创建一个在指定时间后自动取消的context

	//advance.SelectCase() // 随机选择
	//advance.SelectCase2() // 超时处理
	//advance.SelectCase3() // 默认分支
	//advance.SelectCase4() // 非阻塞通信
	//advance.SelectCase5() // 结合context使用
	advance.SelectCase6() // 多路复用
}
