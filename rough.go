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
