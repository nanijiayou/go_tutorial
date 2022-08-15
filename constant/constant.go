package constant

import "fmt"

const (
	BEJING = 0
	SHANGHAI
	SHENZHEN
)

const (
	a = iota
	b
	c
)

const (
	x = 10 * iota
	y
	z
)

const (
	i, j = iota + 1, iota + 2 // iota = 0, i = 0 + 1 = 1, j = 0 + 2 = 2
	k, l                      // iota = 1, k = 1 + 1 = 2, l = 1 + 2 = 3
	m, n                      // iota = 2

	o, p = iota * 2, iota * 3 // iota = 3 o = 3 * 2 = 6, p = 3 * 3 = 9
	q, r                      // iota = 4 q = 4 * 2 = 8, r = 4 * 3 = 12
)

func TestConst() {
	fmt.Println("============= constant ===============")
	const length int = 10
	fmt.Println("length", length)
	// length = 11

	fmt.Println("BEJING = ", BEJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)

	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	fmt.Println("z = ", z)

	fmt.Println("i = ", i, " j = ", j)
	fmt.Println("k = ", k, " l = ", l)
	fmt.Println("m = ", m, " n = ", n)
	fmt.Println("o = ", o, " p = ", p)
	fmt.Println("q = ", q, " r = ", r)
}
