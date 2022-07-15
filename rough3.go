//interface example
package main

import (
	"fmt"
)

type library interface {
	enter_book() books
	book_idadd() uint
}
type book interface {
}
type books struct {
	book_id uint
}

func (b books) enter_book() books {
	return b
}
func (b books) book_idadd() uint {
	return b.book_id + 1
}
func enterdetails_book(lib library) {

	fmt.Println(lib.enter_book())
	fmt.Println(lib.book_idadd())

}

func main() {

	b := books{book_id: 7}
	enterdetails_book(b)
}
