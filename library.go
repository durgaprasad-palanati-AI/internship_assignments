package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//function to check element exist in list
func exist(ids []string, id string) (int, bool) {
	for ix, v := range ids {
		if v == id {
			return ix, true
		}
	}
	return -1, false
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

//book structure
type book struct {
	Book_id     string   `json:"Book_id"`
	Book_name   string   `json:"Book_name"`
	Book_author string   `json:"Book_author"`
	Book_kind   booktype `json:"Book_kind"`
	Pb_copy     bool     `json:"Pb_copy"`
	Db_copy     bool     `json:"Db_copy"`
}

type allbooks []book

var books allbooks

//user structure
type user struct {
	User_id    string   `json:"User_id"`
	User_books []string `json:"User_books"`
	//user_bookcount uint
}
type allusers []user

var users allusers

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createbook(w http.ResponseWriter, r *http.Request) {
	var newbook book
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {

		fmt.Fprintf(w, "enter data ")
	}

	json.Unmarshal(reqBody, &newbook)
	books = append(books, newbook)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newbook)
}

func getOnebook(w http.ResponseWriter, r *http.Request) {
	bookid := mux.Vars(r)["Book_id"]

	for _, singlebook := range books {
		if singlebook.Book_id == bookid {
			json.NewEncoder(w).Encode(singlebook)
		}
	}
}

func getAllbooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func updatebook(w http.ResponseWriter, r *http.Request) {
	bookid := mux.Vars(r)["Book_id"]
	var updatedbook book

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "enter data  to update")
	}
	json.Unmarshal(reqBody, &updatedbook)

	for i, singlebook := range books {
		if singlebook.Book_id == bookid {
			singlebook.Book_name = updatedbook.Book_name
			singlebook.Book_author = updatedbook.Book_author
			singlebook.Book_kind = updatedbook.Book_kind
			singlebook.Pb_copy = updatedbook.Pb_copy
			singlebook.Db_copy = updatedbook.Db_copy
			books = append(books[:i], singlebook)
			json.NewEncoder(w).Encode(singlebook)
		}
	}
}

func deletebook(w http.ResponseWriter, r *http.Request) {
	bookid := mux.Vars(r)["Book_id"]

	for i, singlebook := range books {
		if singlebook.Book_id == bookid {
			books = append(books[:i], books[i+1:]...)
			fmt.Fprintf(w, "The book with ID %v has been deleted successfully", bookid)
		}
	}

}

//new user registration
func registeruser(w http.ResponseWriter, r *http.Request) {
	var newuser user
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {

		fmt.Fprintf(w, "enter data ")
	}

	json.Unmarshal(reqBody, &newuser)
	users = append(users, newuser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newuser)
}

// get users list
func getAllusers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

//get an user by id
func getOneuser(w http.ResponseWriter, r *http.Request) {
	userid := mux.Vars(r)["User_id"]

	for _, singleuser := range users {
		if singleuser.User_id == userid {
			json.NewEncoder(w).Encode(singleuser)
		}
	}
}

// delete a user
func deleteuser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete user")
	userid := mux.Vars(r)["User_id"]
	fmt.Println(userid)
	for i, singleuser := range users {
		if singleuser.User_id == userid {
			users = append(users[:i], users[i+1:]...)
			fmt.Println(users)
			fmt.Fprintf(w, "The user with ID %v has been deleted successfully", userid)
		}
	}

}

//borrow a book
var user_bookids_map = make(map[string][]string) //list of books borrowed by an user
func borrowbook(w http.ResponseWriter, r *http.Request) {
	type bookmap struct {
		Book_id string `json:"Book_id"`
		User_id string `json:"User_id"`
	}

	var borrowbook bookmap
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "enter bookid and userid")
	}
	var book_ids []string
	var user_ids []string
	var User_bookss []string
	json.Unmarshal(reqBody, &borrowbook)
	//fmt.Println(borrowbook)
	//var bx:=0,ux:=0,ubL:=0,ub:=0//book exist,user exists,user book limit,user borrowed
	for _, singlebook := range books {
		book_ids = append(book_ids, singlebook.Book_id)
	}
	for _, singleuser := range users {
		user_ids = append(user_ids, singleuser.User_id)
		User_bookss = append(User_bookss, borrowbook.Book_id)
	}
	//check book id and user id exists or not
	_, bx := exist(book_ids, borrowbook.Book_id) //book index ,book exist(T/F)
	_, ux := exist(user_ids, borrowbook.User_id) //user index ,user exist(T/F)
	if bx {                                      //if book exist in list
		if ux { //if user is in list
			if len(user_bookids_map[borrowbook.User_id]) < 2 {
				_, alx := exist(user_bookids_map[borrowbook.User_id], borrowbook.Book_id) //check book already borrowed
				if alx {
					fmt.Println("user already borrowed same book")
				} else {
					user_bookids_map[borrowbook.User_id] = append(user_bookids_map[borrowbook.User_id], borrowbook.Book_id)
					//singleuser.User_books = append(singleuser.User_books, borrowbook.Book_id)
					//users = append(users, singleuser)
					fmt.Fprintf(w, "The book with ID %v borrowed by %v\n", borrowbook.Book_id, borrowbook.User_id)
					fmt.Fprintf(w, " %v", user_bookids_map)
					json.NewEncoder(w).Encode(user_bookids_map)
				}
			} else {
				fmt.Fprintf(w, "user reached limit to borrow book")
			}
		} else {
			fmt.Fprintf(w, "user id not found")
		}
	} else {
		fmt.Fprintf(w, "book id not found")
	}
}

func main() {
	//initEvents()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/enterbook", createbook).Methods("POST")
	router.HandleFunc("/allbooks", getAllbooks).Methods("GET")
	router.HandleFunc("/book/{Book_id}", getOnebook).Methods("GET")
	router.HandleFunc("/modifybook/{id}", updatebook).Methods("PATCH")
	router.HandleFunc("/deletebook/{Book_id}", deletebook).Methods("DELETE")
	router.HandleFunc("/registeruser", registeruser).Methods("POST")
	router.HandleFunc("/allusers", getAllusers).Methods("GET")
	router.HandleFunc("/user/{User_id}", getOneuser).Methods("GET")
	router.HandleFunc("/deleteuser/{User_id}", deleteuser).Methods("DELETE")
	router.HandleFunc("/borrowbook", borrowbook).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":8080", router))
}
