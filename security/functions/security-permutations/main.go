package main

import "fmt"

func main() {
	var n int
	var m map[int]int // X -> f(x)

	m = make(map[int]int)
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var x int
		fmt.Scan(&x)
		m[i] = x
	}

	for i := 1; i <= n; i++ {
		fmt.Println(m[m[i]])
	}
}
