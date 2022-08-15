package reflectType

import (
	"fmt"
	"reflect"
)

type Pesron struct {
	id   int
	Name string
	age  int
}

func (this Pesron) Show() {
	fmt.Println("person is ", this.Name)
}

func (this Pesron) GetName() string {
	return this.Name
}

func reflectOf(arg interface{}) {
	fmt.Println("type : ", reflect.TypeOf(arg))
	fmt.Println("value : ", reflect.ValueOf(arg))
}

type user struct {
	Name string `info:"name" doc:"your name"`
	Sex  string `info:"sex"`
}

func findTag(info interface{}) {
	t := reflect.TypeOf(info).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagName := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("name = ", tagName, " doc = ", tagDoc)
	}
}

func reflectHandler(input interface{}) {
	fmt.Println("============= reflect ===============")
	// 获取type
	inputType := reflect.TypeOf(input)
	fmt.Println("input type = ", inputType.Name())

	// 获取值
	inputValue := reflect.ValueOf(input)
	fmt.Println("input value = ", inputValue)

	// 获取字段
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i)
		fmt.Printf("%s: %v = %v\n ", field.Name, field.Type, value)
	}

	// 获取方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s %v \n", m.Name, m.Type)
	}
}

func TestRelect() {
	fmt.Println("============= reflect ===============")

	var num float64 = 3.1415926
	reflectOf(num)

	kalai := Pesron{id: 1, Name: "kalai", age: 18}
	kalai.Show()

	reflectHandler(kalai)

	var u user
	findTag(&u)
}
