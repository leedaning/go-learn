package advance

import (
	"errors"
	"fmt"
	"sync"
)

//带错误的Once

type OnceWithError struct {
	once sync.Once
	err  error
}

func (o *OnceWithError) Do(f func() error) error {
	o.once.Do(func() {
		o.err = f()
	})
	return o.err
}

func MySyncOnce6() {
	var once6 OnceWithError
	err := once6.Do(func() error {
		fmt.Printf("Initializing with possible error\n")
		//模拟错误
		return errors.New("initialization failed")
	})

	fmt.Println("Error:", err)

	//再次调用，不会执行函数，但会返回之前的错误
	err = once6.Do(func() error {
		fmt.Println("This won't be executed")
		return nil
	})
	fmt.Println("Error on second call:", err)
}
