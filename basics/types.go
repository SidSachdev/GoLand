package main

import "fmt"

var y = 42
var a = "car"

func main() {
	fmt.Printf("%T\n", y)
	fmt.Println(y)
	//	y = "car"
	//	fmt.Println(y)
	// SINCE GO is STATIC typed this will throw an error
}
