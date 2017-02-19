package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	//Unique N-N mapping variations in Combinatorics is factorial
	all := 1

	for i := 1; i <= n; i++ {
		all *= i
	}

	fmt.Println(all)
}
