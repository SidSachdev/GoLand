package main
import "fmt"
func main() {

	// different types of for statements
	// for statement with a single condition
	x := 1
	for x < 5{
		fmt.Println("Hello World")
		x++
	}

	// another way to do for loop
	y := 1
	for {
		y++
		if y > 10{
			break
		}
		if y % 2!= 0{
			continue
		}
		fmt.Println(y)
	}
	fmt.Println("done")
}
