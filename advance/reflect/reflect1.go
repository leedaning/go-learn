package _reflect

import (
	"fmt"
	"reflect"
)

// 反射
func Reflect1() {

	// 一、获取value、type
	//getValAndType()

	//二、获取和设置值
	//getAndSetVal()

	// 三、检查Value是否有效
	//checkValid()

	// 四、获取类型信息
	getTypeInfo()
}

// 一、获取value、type
func getValAndType() {

	var x float64 = 3.4

	// 获取value
	v := reflect.ValueOf(x)
	//获取type
	t := reflect.TypeOf(x)

	fmt.Printf("x is of type %T, t is :%v, v is :%v \n", x, v, t)

	//从value中获取Type
	fmt.Printf("Type from value :%v \n", v.Type())

	//Kind表示底层类型
	fmt.Printf("Kind:%v \n", v.Kind())
}

// 二、获取和设置值
func getAndSetVal() {
	var x float64 = 3.4

	p := reflect.ValueOf(&x) //修改值需要传入指针
	vp := p.Elem()           // 获取指针指向的值
	vp.SetFloat(3.1415)
	fmt.Printf("x is of type %T, p type is :%v \n", x, p.Type())

}

// 三、检查Value是否有效
func checkValid() {

	var x float64 = 3.4

	// 获取value
	v := reflect.ValueOf(x)

	if v.IsValid() {
		fmt.Printf("x is valid\n")
	} else {
		fmt.Printf("x is invalid\n")
	}

	var y interface{}
	//y = 3.141
	v = reflect.ValueOf(y)
	if v.IsValid() {
		fmt.Printf("y is valid\n")
	} else {
		fmt.Printf("y is invalid\n")
	}
}

type MyStruct struct {
	Field1 int
	Field2 string
}

// 四、获取类型信息
func getTypeInfo() {
	var s MyStruct

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	fmt.Printf("Name:%v\n", t.Name())      // MyStruct
	fmt.Printf("Kind:%v\n", t.Kind())      // struce
	fmt.Println("NumField:", t.NumField()) // 2
	fmt.Printf("v is :%v \n", v)           // {0 }

	// 遍历结构体字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("%d. %s (%s)\n", i, field.Name, field.Type)
	}
}
