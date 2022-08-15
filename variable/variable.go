package variable

import "fmt"

// 全局变量
var gloA = 100
var gloB = "abc"

// := 只能在函数体内声明
// gloC := "gloC"

func TestVar() {
	fmt.Println("============= variable ===============")
	var a int
	fmt.Println("a = ", a)

	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n ", b)

	var c = "hello"
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	g := 3.14
	fmt.Println("e = ", g)
	fmt.Printf("type of g = %T\n", g)

	fmt.Println("gloA = ", gloA)
	fmt.Println("gloB = ", gloB)

	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, "yy = ", yy)

	var zz, ww = 100, "abc"
	fmt.Println("zz = ", zz, "ww = ", ww)

	var (
		vv int = 100
		jj int = 200
	)
	fmt.Println("vv = ", vv, "jj = ", jj)
}
