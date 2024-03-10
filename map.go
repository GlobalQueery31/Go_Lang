package main

import "fmt"

func main() {
	mp := make(map[int]int)

	mp[0] = 69
	mp[1] = 79

	fmt.Println(mp[1])

	delete(mp, 1)

	fmt.Println(mp[1])
	// v := []int {1,1,2,3,4,5,5,5}

	// for i:=0; i<v.Len(); i++ {
	// 	mp[v[i]]++;
	// }

	// for i:=0; i<mp.Len(); i++ {
	// 	fmt.Println(mp[i])
	// }
}
