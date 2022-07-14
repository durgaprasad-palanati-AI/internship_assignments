package main

import (
	"fmt"
	"os"
)

//function to check element exist in list
func exist(ids []uint, id uint) (int, bool) {
	for ix, v := range ids {
		if v == id {
			return ix, true
		}
	}
	return -1, false
}

//remove element from array
func del_element(a []uint, index int) []uint {
	return append(a[:index], a[index+1:]...)
}

//
func main() {
	var book_id uint                                  //variable to store book's id
	var user_id uint                                  //variable to store user's id
	type add_book func(book_id uint)                  //function to add new book
	type add_user func(user_id uint)                  //function to add new user
	type borrow_book func(book_id uint, user_id uint) //function to borrow a book by user
	type return_book func(book_id uint, user_id uint) //function to return a book by user
	var book_ids []uint
	var user_ids []uint
	var book_user_map = make(map[uint][]uint)    //list of books borrowed by each user
	var user_bookids_map = make(map[uint][]uint) //list of books borrowed by user
	/*library type*/
	type Library struct {
		add_newbook    add_book    //function of type add_book
		books_count    uint        // variable to store the number of books
		add_newuser    add_user    //function of type add_user
		users_count    uint        // variable to store the number of users
		borrow_newbook borrow_book //function of type borrow_book
		return_mybook  return_book //function of type return_book
	}
	/*method to add new book*/
	new_bookentry := Library{
		add_newbook: func(book_id uint) {
			fmt.Println("Enter new book_id")
			fmt.Scanf("%d\n", &book_id) //enter book id & store it in book_id
			fmt.Println("book added with id=", book_id)
			book_ids = append(book_ids, book_id) //add the new book in a list
		}}
	/*method to add new user*/
	new_userentry := Library{
		add_newuser: func(user_id uint) {
			fmt.Println("Enter new user_id")
			fmt.Scanf("%d\n", &user_id) //enter user id & store it in user_id
			fmt.Println("user added with id=", user_id)
			user_ids = append(user_ids, user_id) //add the new user in a list
		}}
	/*method to borrow a book*/
	new_borrowbook := Library{
		borrow_newbook: func(book_id uint, user_id uint) {
			fmt.Println("Enter book_id to borrow")
			fmt.Scanf("%d\n", &book_id)
			fmt.Println("Enter your user_id to borrow a book")
			fmt.Scanf("%d\n", &user_id)
			//check book id and user id exists or not
			_, bx := exist(book_ids, book_id) //book index ,book exist(T/F)
			_, ux := exist(user_ids, user_id) //user index ,user exist(T/F)
			if bx {
				if ux {
					if len(user_bookids_map[user_id]) < 2 {
						user_bookids_map[user_id] = append(user_bookids_map[user_id], book_id)
						fmt.Println(user_id, " borrowed book with id=", book_id)
					} else {
						fmt.Println(user_id, " reached limit to borrow book")
					}
				} else {
					fmt.Println("user id not found")
				}
			} else {
				fmt.Println("book with", book_id, " id not found")
			}
			book_user_map[user_id] = user_bookids_map[user_id]
		}}
	/*method to return a book*/
	book_return := Library{
		return_mybook: func(book_id uint, user_id uint) {
			fmt.Println("Enter book_id to return")
			fmt.Scanf("%d\n", &book_id)
			fmt.Println("Enter your user_id")
			fmt.Scanf("%d\n", &user_id)
			//check book id and user id exists or not
			//_, bx := exist(book_ids, book_id) //book index ,book exist(T/F)
			_, ux := exist(user_ids, user_id) //user index ,user exist(T/F)
			if ux {
				if len(user_bookids_map[user_id]) > 0 {
					ubx, uxb := exist(user_bookids_map[user_id], book_id) //user's book index ,user's book exist(T/F)
					if uxb {
						//delete book from user's book list
						user_bookids_map[user_id] = del_element(user_bookids_map[user_id], ubx)
						fmt.Println(user_id, " returned book with id=", book_id)
					} else {
						fmt.Println("book with", book_id, " id not found")
					}
				} else {
					fmt.Println(user_id, " not borrowed any books")
				}
			} else {
				fmt.Println("user id not found")

			}
			book_user_map[user_id] = user_bookids_map[user_id]
		}}
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
			new_bookentry.add_newbook(book_id)     //calling method to add a new book
			fmt.Println("book ids are:", book_ids) //print book ids
		case 2:
			new_userentry.add_newuser(user_id)     //calling method to add a new book
			fmt.Println("user ids are:", user_ids) //print user ids
		case 3:
			new_borrowbook.borrow_newbook(book_id, user_id) //calling method to add a new book
		case 4:
			//print user and book mapping
			if len(book_user_map) == 0 {
				fmt.Println("No books borrowed")
			} else {
				for key, value := range book_user_map {
					fmt.Printf("user with id=%d borrowed book with id= %v\n", key, value)
				}
			}
		case 5:
			//call return book function
			book_return.return_mybook(book_id, user_id)
		case 6:
			//Exits
			os.Exit(4)
		}
	}
}
