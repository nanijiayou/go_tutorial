package pointers

import "fmt"

func changeValue(p int) {
	p = 10
}

func changeValue2(p *int) {
	*p = 10
}

func swap(p1 *int, p2 *int) {
	var temp int
	temp = *p1
	*p1 = *p2
	*p2 = temp
}

func TestPointers() {
	fmt.Println("============= pointers ===============")
	a := 1
	changeValue(a)
	fmt.Println("a = ", a)

	changeValue2(&a)
	fmt.Println("a = ", a)

	c := 2
	d := 4
	swap(&c, &d)
	fmt.Println("c = ", c, "d = ", d)
}
