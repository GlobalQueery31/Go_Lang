package main

import (
	"fmt"
	"sync"
)

func increament(a *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	*a++
}

func decreament(a *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	*a++
}

func main() {
	a:=8
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(1)
	go increament(&a,&mu,&wg)
	wg.Add(1)
	go decreament(&a,&mu,&wg)
	fmt.Println(a)
	wg.Wait()

}
