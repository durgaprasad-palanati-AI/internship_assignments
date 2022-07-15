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

//enum
type booktype string

const (
	ebook        booktype = "ebook"
	Audiobook    booktype = "Audiobook"
	Hardback     booktype = "Hardback"
	Paperback    booktype = "Paperback"
	Encyclopedia booktype = "Encyclopedia"
	Magazine     booktype = "Magazine"
	Comic        booktype = "Comic"
)

//book interface
type book interface {
	kind_book() booktype
	name_book() string
	author_book() string
	id_book() uint
}

//book structure
type books struct {
	book_id     uint
	book_name   string
	book_author string
	book_kind   booktype
	pb_copy     bool
	db_copy     bool
}

// function to return kind of book
func (b books) kind_book() booktype {
	return b.book_kind
}

// function to return name of book
func (b books) name_book() string {
	return b.book_name
}

// function to return author of book
func (b books) author_book() string {
	return b.book_author
}

// function to return id of book
func (b books) id_book() uint {
	return b.book_id
}

type borrow func(uint) bool
type Phy_books struct {
	books
	//borrow_Phybook borrow
}
type Digi_books struct {
	books
	//borrow_Digibook borrow
}

//constructor
func New_phybook(b books) *Phy_books {
	pb := new(Phy_books)
	pb.books = b
	return pb
}

//constructor
func New_digibook(b books) *Digi_books {
	db := new(Digi_books)
	db.books = b
	return db
}

//user structure
type users struct {
	user_id        uint
	user_bookcount uint
}

var b books                                       // variable of book type
var u users                                       //variable of user type
var user_id uint                                  //variable to store user's id
type add_book func(b books)                       //function to add new book
type add_user func(u users)                       //function to add new user
type borrow_book func(book_id uint, user_id uint) //function to borrow a book by user
type return_book func(book_id uint, user_id uint) //function to return a book by user
type library struct {
	add_newbook    add_book    //function of type add_book
	add_newuser    add_user    //function of type add_user
	borrow_newbook borrow_book //function of type borrow_book
	return_mybook  return_book //function of type return_book
}

// function to return all details of book by implementing book interface
func getbook_details(b book) {
	fmt.Println("book details are:", b)
	/*
		fmt.Println(b.id_book())
		fmt.Println(b.name_book())
		fmt.Println(b.author_book())
		fmt.Println(b.kind_book())
	*/
}
func main() {
	var books_list []books
	var users_list []users
	var book_ids []uint
	var user_ids []uint
	var book_user_map = make(map[uint][]uint)    //list of books borrowed by each user
	var user_bookids_map = make(map[uint][]uint) //list of books borrowed by an user
	/*method to add new book*/
	new_bookentry := library{
		add_newbook: func(b books) {
			fmt.Println("enter id of book")
			fmt.Scanf("%d\n", &b.book_id)
			fmt.Println("enter name of book")
			fmt.Scanf("%s\n", &b.book_name)
			fmt.Println("enter author of book")
			fmt.Scanf("%s\n", &b.book_author)
			fmt.Println("enter type of book")
			fmt.Scanf("%s\n", &b.book_kind)
			fmt.Println("Is phy_copy book available?")
			fmt.Scanf("%t\n", &b.pb_copy)
			fmt.Println("Is digi_copy book available?")
			fmt.Scanf("%t\n", &b.db_copy)
			fmt.Println("book added with id=", b.book_id)
			books_list = append(books_list, b)     //add the new book in a list
			book_ids = append(book_ids, b.book_id) //add the new book in a list
		}}
	/*method to add new user*/
	new_userentry := library{
		add_newuser: func(u users) {
			fmt.Println("Enter new user_id")
			fmt.Scanf("%d\n", &u.user_id)          //enter user id & store it in user_id
			user_ids = append(user_ids, u.user_id) //add the new user in a list
			users_list = append(users_list, u)     //add the new user in a list
		}}

	/*method to borrow a book*/
	new_borrowbook := library{
		borrow_newbook: func(book_id uint, user_id uint) {
			fmt.Println("Enter book_id to borrow")
			fmt.Scanf("%d\n", &book_id)
			fmt.Println("Enter your user_id to borrow a book")
			fmt.Scanf("%d\n", &user_id)

			//check book id and user id exists or not
			_, bx := exist(book_ids, book_id) //book index ,book exist(T/F)
			_, ux := exist(user_ids, user_id) //user index ,user exist(T/F)
			if bx {                           //if book exist in list
				if ux { //if user is in list
					//check the limit of user boorowed books
					if len(user_bookids_map[user_id]) < 2 {
						_, alx := exist(user_bookids_map[user_id], book_id) //check book already borrowed by user
						if alx {
							fmt.Println("user already borrowed same book")
						} else {
						copyLoop:
							for true {
								fmt.Println("enter your choice")
								c := 0
								fmt.Println("1.Physical copy")
								fmt.Println("2.Digital copy")
								fmt.Scanf("%d\n", &c)
								switch c {
								case 1:
									for _, v := range books_list {
										if v.pb_copy { //check whether physical copy exist
											if v.book_id == book_id { //check book id exist or not
												// add the book to existing user's book list
												user_bookids_map[user_id] = append(user_bookids_map[user_id], book_id)
												fmt.Println("user-", user_id, " borrowed book with id=", book_id)
											} else {
												fmt.Println("Physical copy is not available")
											}
										}
									}
									break copyLoop
								case 2:
									for _, v := range books_list {
										if v.db_copy { //check whether digital copy exist
											if v.book_id == book_id { //check book id exist or not
												// add the book to existing user's book list
												user_bookids_map[user_id] = append(user_bookids_map[user_id], book_id)
												fmt.Println("user-", user_id, " borrowed book with id=", book_id)
											} else {
												fmt.Println("Digital copy is not available")
											}
										}
									}
									break copyLoop //come out of inner
								}
							}

						}
					} else {
						fmt.Println("user-", user_id, " reached limit to borrow book")
					}
				} else {
					fmt.Println("user id-", user_id, " not found")
				}
			} else {
				fmt.Println("book with", book_id, " id not found")
			}
			book_user_map[user_id] = user_bookids_map[user_id]
		}}
	/*method to return a book*/
	book_return := library{
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
						fmt.Println("book with", book_id, " id not found for userid-", user_id)
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
		fmt.Println("1.add book")
		fmt.Println("2.enter new user")
		fmt.Println("3.borrow a book")
		fmt.Println("4.user & book map")
		fmt.Println("5.return a book")
		fmt.Println("6.get all books")
		fmt.Println("7.Get physical books list")
		fmt.Println("8.Get digital books list")
		fmt.Println("9.Get books with both physical and digital copy")
		fmt.Println("10.EXIT")
		fmt.Scanf("%d\n", &i)
		switch i {
		case 1:
			new_bookentry.add_newbook(b)              //calling method to add a new book
			fmt.Println("list of books:", books_list) //print book ids
		case 2:
			new_userentry.add_newuser(u)
			fmt.Println("user's list is:", users_list) //print user ids
		case 3:
			new_borrowbook.borrow_newbook(b.book_id, user_id) //calling method to add a new book
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
			book_return.return_mybook(b.book_id, user_id)
		case 6:
			fmt.Println("books are")
			for _, v := range books_list {
				getbook_details(v)
			}
		case 7:
			fmt.Println("Physical books are")
			for _, v := range books_list {
				if v.pb_copy {
					fmt.Println(*New_phybook(v))
				}
			}
		case 8:
			fmt.Println("Digital books are")
			for _, v := range books_list {
				if v.db_copy {
					fmt.Println(*New_digibook(v))
				}
			}
		case 9:
			var pd_books []books
			for _, v := range books_list {
				if v.db_copy && v.pb_copy {
					pd_books = append(pd_books, v)
				}
			}
			if len(pd_books) > 0 {
				fmt.Println("books with both Physical and digital copy")
				fmt.Println(pd_books)
			} else {
				fmt.Println("No books with both Physical and digital copy")
			}
		case 10:
			os.Exit(4)
		}
	}
}
