package functions

import "fmt"

func TestFunc(a string, b int) int {
	fmt.Println("a = ", a, "b = ", b)
	c := 10
	return c
}

func TestFuncMultiReturn(a string, b int) (int, int) {
	fmt.Println("a = ", a, "b = ", b)
	return 1, 2
}

func TestFuncMultiReturnWithName(a string, b int) (r1 int, r2 int) {
	fmt.Println("a = ", a, "b = ", b)
	r1 = 3
	r2 = 4
	return
}

func TestFuncs() {
	fmt.Println("============= functions ===============")
	r1 := TestFunc("ab", 1)
	fmt.Println("res1 = ", r1)

	r2, r3 := TestFuncMultiReturn("ef", 2)
	fmt.Println("res2 = ", r2, "res3 = ", r3)

	r4, r5 := TestFuncMultiReturnWithName("ef", 2)
	fmt.Println("res4 = ", r4, "res5 = ", r5)
}
