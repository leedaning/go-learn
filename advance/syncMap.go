package advance

import (
	"fmt"
	"sync"
)

func MySyncMap() {
	var m sync.Map

	//存储键值对
	m.Store("name", "Leen")
	m.Store("age", 18)

	// 读取值
	if v, ok := m.Load("name"); ok {
		fmt.Println("name:", v.(string))
	}

	if v, ok := m.Load("age"); ok {
		fmt.Println("age:", v.(int))
	}

	// 遍历所有键值对
	m.Range(func(key, value interface{}) bool {
		fmt.Println("key:", key, "value:", value)
		return true
	})

	name, res := m.LoadOrStore("names", "daning")
	fmt.Println("names:", name, "res:", res)

	//删除键
	m.Delete("age")

	// 遍历所有键值对
	m.Range(func(key, value interface{}) bool {
		fmt.Println("key:", key, "value:", value)
		return true
	})
}
