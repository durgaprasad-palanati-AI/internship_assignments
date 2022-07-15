package main

import (
	"fmt"
	"os/user"
)

type booktype int

const (
	ebook booktype = iota + 1
	Audiobook
	Hardback
	Paperback
	Encyclopedia
	Magazine
	Comic
)

func (bt booktype) String() string {
	return [...]string{"ebook", "Audiobook", "Hardback", "Paperback", "Encyclopedia", "Magazine", "Comic"}[bt]
}

// interface
type library struct {
	enterdetails_book() books
	enterdetails_user() users
}
type book interface {
	getdetails_book() books
	borrow_map(book_id uint, user_id uint) bool
}

//structures
type books struct {
	book_id     uint
	book_author string
	book_name   string
	book_type   string
}
type users struct {
	user_id   uint
	user_name string
}

func (b books) enterdetails_book() books {
	return b
}
func enter_bookdetails(b books) books {
	return b.enterdetails_book()
}
func (u users) enterdetails_user() users {
	return u
}
func enter_userdetails(l library) users {
	return l.enterdetails_user()
}
func (b books) borrow_map() bool {
	return true
}
func borrow_bookmap(b book) {
	var bk books
	var u users
	fmt.Println("borrowed book", b.borrow_map(bk.book_id, u.user_id))
}
func main() {
	var b books
	var u users
	for true {
		i := 0
		fmt.Println("1.newbook entry")
		fmt.Println("2.newuser entry")
		fmt.Println("3.borrow a book")
		fmt.Println("4.user & book map")
		fmt.Println("5.return a book")
		fmt.Println("6.EXIT")
		fmt.Scanf("%d\n", &i)

		switch i {
		case 1:
			fmt.Println("enter id of book")
			fmt.Scanf("%d\n", &b.book_id)
			fmt.Println("enter name of book")
			fmt.Scanf("%s\n", &b.book_name)
			fmt.Println("enter author of book")
			fmt.Scanf("%s\n", &b.book_author)
			fmt.Println("enter type of book")
			fmt.Scanf("%s\n", &b.book_type)
			bks := enter_bookdetails(b)
			fmt.Println("details of book", bks)
		case 2:
			fmt.Println("enter id of user")
			fmt.Scanf("%d\n", &u.user_id)
			fmt.Println("enter name of user")
			fmt.Scanf("%s\n", &u.user_name)
		case 3:
			fmt.Println("enter id of book")
			fmt.Scanf("%d\n", &b.book_id)
			fmt.Println("enter id of user")
			fmt.Scanf("%d\n", &u.user_id)
			borrow_bookmap(b.book_id,u.user_id)
		}
	}

}
