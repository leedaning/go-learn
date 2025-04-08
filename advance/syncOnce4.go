package advance

import (
	"fmt"
	"sync"
)

type Singleton4 struct {
	name string
}

var (
	instance4 *Singleton4
	once4     sync.Once
)

func Instance4() *Singleton4 {
	once4.Do(func() {
		//下面这两种方式等价
		/*instance4 = new(Singleton4)
		instance4.name = "Leen"*/

		instance4 = &Singleton4{name: "Leen"}
		fmt.Printf("生命单例模式模式\n")
	})
	return instance4
}

func MySyncOnce4() {

	s1 := Instance4()
	s2 := Instance4()

	fmt.Printf("s1 name:%s\n", s1.name)
	fmt.Printf("s2 name:%s\n", s2.name)

	fmt.Println(s1 == s2)

}
