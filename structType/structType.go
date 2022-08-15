package structType

import "fmt"

type myint int

type Book struct {
	name string
	auth string
}

func update(book Book) {
	book.name = "update c++"
}

func updateBook(book *Book) {
	book.name = "update c++"
}

func TestStruct() {
	fmt.Println("============= struct ===============")
	var a myint = 10
	fmt.Printf("type of a is %T\n", a)

	var book Book
	book.name = "c++"
	book.auth = "kalai"
	fmt.Printf("book is %v\n", book)

	update(book)
	fmt.Printf("book is %v\n", book)

	updateBook(&book)
	fmt.Printf("book is %v\n", book)
}
