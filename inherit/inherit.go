package inherit

import (
	"fmt"
)

type Human struct {
	name string
	age  int
}

func (this *Human) Eat() {
	fmt.Println("name = ", this.name, "is eating")
}

func (this *Human) Walking() {
	fmt.Println("name = ", this.name, "is walking")
}

type SuperMan struct {
	Human
	speed int
}

func (this *SuperMan) Eat() {
	fmt.Println("super man name = ", this.name, "is eating")
}

func (this *SuperMan) Fly() {
	fmt.Println("super man name = ", this.name, "is flying")
}

func TestInherit() {
	fmt.Println("============= inherit ===============")
	kalai := Human{name: "kalai", age: 18}
	kalai.Walking()

	jn := SuperMan{Human{name: "super", age: 10000}, 1000}
	jn.Eat()
	jn.Walking()
	jn.Fly()

	var sp SuperMan
	sp.name = "super kalai"
	sp.age = 1000
	sp.speed = 100000
	sp.Eat()
	sp.Walking()
	sp.Fly()
}
