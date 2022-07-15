/*package main

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
*/
//return values example
/*
package main

import (
	"fmt"
)

func vals() (int, int) {
	return 3, 7
}
func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	//fmt.Println(reflect.TypeOf(vals()))

	_, c := vals()
	fmt.Println(c)
}
*/
//remove array element
/*
package main

import (
	"fmt"
)

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {
	all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("all: ", all) //[0 1 2 3 4 5 6 7 8 9]
	all = RemoveIndex(all, 5)

	//fmt.Println("all: ", all)                 //[0 1 2 3 4 6 7 8 9 9]
	fmt.Println("removeIndex: ", all) //[0 1 2 3 4 6 7 8 9]
}
*/
//interface example
/*
package main

import "fmt"

// interface
type Shape interface {
	area() float32
}

// struct to implement interface
type Rectangle struct {
	length, breadth float32
}

// use struct to implement area() of interface
func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

// access method of the interface
func calculate(s Shape) {
	fmt.Println("Area:", s.area())
}

// main function
func main() {

	// assigns value to struct members
	rect := Rectangle{7, 4}

	// call calculate() with struct variable rect
	calculate(rect)

} */
package main

import "fmt"

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

type books struct {
	book_author string
	book_name   string
	book_type   string
}

// interface
type book interface {
	details_Book()
}

func (b books) details_Book() {
	fmt.Println("details of book", b)
}
func print_bookdetails() {
	var bt booktype
	fmt.Println("enter type of book")
	fmt.Scanf("%d\n", &bt)
	var b books
	switch bt {
	case ebook:
		b = books{"durga", "GOLANG", "ebook"}
	case Audiobook:
		b = books{"durga", "GOLANG", "Audiobook"}
	case Hardback:
		b = books{"durga", "GOLANG", "Hardback"}
	case Paperback:
		b = books{"durga", "GOLANG", "Paperback"}
	case Encyclopedia:
		b = books{"durga", "GOLANG", "Encyclopedia"}
	case Magazine:
		b = books{"durga", "GOLANG", "Magazine"}
	case Comic:
		b = books{"durga", "GOLANG", "Comic"}
	}

	b.details_Book()
}
func main() {
	print_bookdetails()
}
