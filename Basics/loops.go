package main

import "fmt"

func main() {
	for i:=0;i<11;i+=1{
		fmt.Print(i," ")
	}

	fmt.Println("\n")
	i:=0

	for i<11{
		fmt.Println(i)
		i++;
	}
}
