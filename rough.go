package main

import (
	"fmt"
)

func main() {

	type books struct {
		ids   uint
		names []string
		rolls []uint
		marks []uint
	}
	book := books{ids: 123, names: []string{"abc", "bcd"}, rolls: []uint{1, 2},
		marks: []uint{100, 200}}
	fmt.Println(book.ids)
	fmt.Println(book.names)
	fmt.Println(book.rolls)

	fmt.Println(book.marks[0])
	book.marks[0] = 300
	fmt.Println(book.marks[0])
}
