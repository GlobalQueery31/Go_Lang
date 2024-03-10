package main

import (
	"fmt"
	"sync"
)

func printnum(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		//time.Sleep(100 * time.Millisecond)
	}
}

func printchar(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c\n",i)
		//time.Sleep(100 * time.Millisecond)
	}
}
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go printnum(&wg)

	wg.Add(1)
	go printchar(&wg)

	wg.Wait()
	// time.Sleep(1500 * time.Millisecond)
}
