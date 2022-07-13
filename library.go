package main

import (
	"fmt"
)

func main() {
	var book_id uint                         //variable to store book's id
	var user_id uint                         //variable to store user's id
	type add_book func(book_id uint)         //function to add new book
	type add_user func(user_id uint)         //function to add new user
	type borrow_book func(book_id uint) bool ////function to boorow a book by user

	/*library type*/
	type Library struct {
		add_newbook    add_book    //function of type add_book
		books_count    uint        // variable to store the number of books
		add_newuser    add_user    //function of type add_user
		users_count    uint        // variable to store the number of users
		borrow_newbook borrow_book //function of type borrow_book
	}
	//method to add new book
	new_bookentry := Library{
		add_newbook: func(book_id uint) {
			fmt.Println("Enter new book_id")
			fmt.Scanf("%d\n", &book_id) //enter book id & store it in book_id
			fmt.Println("book added with id=", book_id)
		}}

	new_bookentry.add_newbook(book_id) //calling method to add a new book
	//method to add new user
	new_userentry := Library{
		add_newuser: func(user_id uint) {
			fmt.Println("Enter new user_id")
			fmt.Scanf("%d\n", &user_id) //enter user id & store it in user_id
			fmt.Println("user added with id=", user_id)
		}}
	new_userentry.add_newuser(user_id) //calling method to add a new book
	//method to borrow a book
	new_borrowbook := Library{
		borrow_newbook: func(book_id uint) bool {
			fmt.Println("Enter new book_id to borrow")
			fmt.Scanf("%d\n", &book_id) //enter user id & store it in user_id
			fmt.Println("user borrowed book with id=", book_id)
			return true
		}}
	new_borrowbook.borrow_newbook(book_id) //calling method to add a new book
}
