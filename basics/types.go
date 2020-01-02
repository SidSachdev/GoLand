package main

import "fmt"

var y = 42
var a = "car"

// define new type
type catname string

func main() {
	fmt.Printf("%T\n", a)
	fmt.Println(a)
	//	y = "car"
	//	fmt.Println(y)
	// SINCE GO is STATIC typed this will throw an error

	// working on newly created type
	name := catname("Maya")
	fmt.Println(name)
	fmt.Printf("%T\n", name)

	// a = name
	// WILL NOT WORK

	//conversion NOT casting
	// lets convert catname to string
	a = string(name)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
