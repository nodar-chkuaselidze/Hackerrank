package main

import "fmt"

func main() {
	//let's inverse
	var n int
	var mappings map[int]int // F(x) -> y, save y : x

	fmt.Scan(&n)
	mappings = make(map[int]int)

	for i := 1; i <= n; i++ {
		var x int

		fmt.Scan(&x)

		mappings[x] = i
	}

	for i := 1; i <= n; i++ {
		fmt.Println(mappings[i])
	}
}
