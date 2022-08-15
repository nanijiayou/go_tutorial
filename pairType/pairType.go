package pairType

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a book")
}

func TestPairType() {
	fmt.Println("============= pair ===============")

	var a string
	a = "abced"

	var allType interface{}
	allType = a

	value, _ := allType.(string)
	fmt.Println(value)

	b := &Book{}

	var r Reader
	r = b
	r.ReadBook()

	var w Writer
	w = r.(Writer)
	w.WriteBook()
}
