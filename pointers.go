package main

import "fmt"

func main() {
	var value int = 1000
	var pointer *int = &value
	println(value)                //1000
	println(pointer)              //0xfffffffff
	println(*pointer)             //1000
	(*pointer)++		  			      //1001
	*pointer = *pointer + 10	    //1011
	fmt.Println(*pointer)			        //1011
	println(*pointer + *pointer)  //1011 + 1011 = 2022

}
