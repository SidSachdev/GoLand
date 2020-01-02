package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

type count int

var a count

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	a = 42
	fmt.Println(a)

}
