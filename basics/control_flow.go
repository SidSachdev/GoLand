package main

import "fmt"

func main() {
	// For loop in Golang
	// for init; condition; post {}
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
