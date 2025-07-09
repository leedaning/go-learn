package basic

import "fmt"

func MyPointer() {
	x := 10
	p := &x
	fmt.Printf("x is %v, p is %v\n", x, *p)
	fmt.Printf("x is of type %T, x add :%v\n", x, &x)
	*p = 20
	fmt.Printf("x is %v, p is %v\n", x, *p)
}
