package advance

import (
	"fmt"
	"sync"
)

// 延迟初始化
type ExpensiveResource struct {
	//假设这是一个初始化成本很高的资源
}

var (
	resource *ExpensiveResource
	once5    sync.Once
)

func GetResource() *ExpensiveResource {
	once5.Do(func() {
		fmt.Printf("Initializing expensive resource ...\n")
		resource = &ExpensiveResource{}
	})
	return resource
}

func MySyncOnce5() {
	//只有在第一次调用时才会初始化
	_ = GetResource()
	_ = GetResource() // 不会再次初始化
}
