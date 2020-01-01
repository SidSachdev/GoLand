package main

import "fmt"

func main() {
	x := 32
	// first time declaration :=
	fmt.Println("Hello World", x)
	// in future, once a var is declared, use =
	x = 99
	fmt.Println("Hello World", x)

	// example of operators
	y := 100 + 24
	fmt.Println("Hello World", y)
}
