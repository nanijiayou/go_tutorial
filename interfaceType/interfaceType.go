package interfaceType

import (
	"fmt"
)

func foo(arg interface{}) {
	fmt.Println("foo func is called")
	fmt.Println(arg)

	// 断言类型
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string")
	} else {
		fmt.Println("arg is ", value)
		fmt.Printf("value type is %T\n", value)
	}
}

type Book struct {
	name string
}

func TestInterFace() {
	fmt.Println("============= interface ===============")

	boo := Book{name: "kalai"}

	foo(boo)
	foo("name")
	foo(100)
}
