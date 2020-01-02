package main

import "fmt"

var x = 42

func main() {
	fmt.Printf("%T\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%#x\n", x)

	s := fmt.Sprintln("car")
	fmt.Println(s)

}
