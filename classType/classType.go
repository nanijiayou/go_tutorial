package classType

import "fmt"

type Person struct {
	name string
	age  int
}

func (this *Person) ShowName() {
	fmt.Println("Name = ", this.name)
}

func (this *Person) GetName() string {
	return this.name
}

func (this *Person) SetName(name string) {
	this.name = name
}

func TestClass() {
	fmt.Println("============= class ===============")
	kalai := Person{name: "kalai", age: 18}
	kalai.ShowName()

	kalai.SetName("laikao")
	kalai.ShowName()
}
