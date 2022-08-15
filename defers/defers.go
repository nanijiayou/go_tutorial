package defers

import "fmt"

func deferFunc() int {
	fmt.Println("defer func called")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func TestDefers() {
	fmt.Println("============= defers ===============")
	defer fmt.Println("test defers func end")

	fmt.Println("test defers")

	returnAndDefer()
}
