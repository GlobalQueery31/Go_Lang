package main

import "fmt"

func main() {
	var i int
	fmt.Println("Enter a number")
	fmt.Scanln(&i)
	switch i {
	case 0:
		fmt.Println("This is ", i)
	case 1:
		fmt.Println("This is ", i)
	case 2:
		fmt.Println("This is ", i)
	default:
		fmt.Println("Ghar chala ja")
	}

}
