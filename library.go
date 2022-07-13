package main

import "fmt"

func main() {
	type add_book func(book_id uint) uint
	type Library struct {
		add_newbook add_book
	}
	books := Library{add_newbook: func(book_id uint) uint { return book_id }}
	fmt.Println("book added with id=", books.add_newbook(123))
}
