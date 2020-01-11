package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("This shouldn't print")
	case (2 == 4):
		fmt.Println("This shouldn't print")
	case (3 == 6):
		fmt.Println("prints")
	default:
		fmt.Println("Default")
	}

	// more different switch  with expr
	switch "car" {
	case "car":
		fmt.Println("print")
	case "fly":
		fmt.Println("This shouldn't print")
	default:
		fmt.Println("Default")
	}
}
