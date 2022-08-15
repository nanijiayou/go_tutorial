package polymorphism

import (
	"fmt"
)

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("cat is sleeping")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("dog is sleeping")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "dog"
}

func showAnimalColor(animal AnimalIF) {
	fmt.Println("animal color = ", animal.GetColor())
}

func TestPolymorphism() {
	fmt.Println("============= polymorphism ===============")

	var animal AnimalIF
	animal = &Cat{"yellow"}

	animal.Sleep()
	showAnimalColor(animal)

	animal = &Dog{"black"}
	animal.Sleep()
	showAnimalColor(animal)
}
