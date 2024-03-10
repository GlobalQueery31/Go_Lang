package main

import "fmt"

func sum(a,b int , ch chan int){
	s:=0
	for i:=a;i<=b;i++{
		s+=i;
	}
	ch <- s
}

func main() {
	ch := make(chan int)
	go sum(1,100,ch)
	output := <- ch
	fmt.Println(output)
}
