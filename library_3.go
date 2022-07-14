package main

import (
	"fmt"
)

//function to check element exist in list
func exist(ids []uint, id uint) bool {
	for _, v := range ids {
		if v == id {
			return true
		}
	}
	return false
}
func main() {
	var book_id uint                                  //variable to store book's id
	var user_id uint                                  //variable to store user's id
	type add_book func(book_id uint)                  //function to add new book
	type add_user func(user_id uint)                  //function to add new user
	type borrow_book func(book_id uint, user_id uint) ////function to boorow a book by user
	var book_ids []uint
	var user_ids []uint
	/*type user_bookids_map map[uint][]uint
	type book_user_map map[uint][]uint*/

	/*library type*/
	type Library struct {
		add_newbook    add_book    //function of type add_book
		books_count    uint        // variable to store the number of books
		add_newuser    add_user    //function of type add_user
		users_count    uint        // variable to store the number of users
		borrow_newbook borrow_book //function of type borrow_book

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
	var book_user_map = make(map[uint][]uint)
	var user_bookids_map = make(map[uint][]uint)
	/*method to borrow a book*/
	new_borrowbook := Library{
		borrow_newbook: func(book_id uint, user_id uint) {
			fmt.Println("Enter book_id to borrow")
			fmt.Scanf("%d\n", &book_id)
			fmt.Println("Enter your user_id to borrow a book")
			fmt.Scanf("%d\n", &user_id)

			//check book id and user id exists or not
			if exist(book_ids, book_id) {
				if exist(user_ids, user_id) {
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
			for key, value := range book_user_map {
				fmt.Printf("user with id=%d borrowed book with id= %v\n", key, value)
			}

		}}
	for true {
		i := 0
		fmt.Println("1.newbook entry")
		fmt.Println("2.newuser entry")
		fmt.Println("3.borrow a book")
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
		}
	}
}
