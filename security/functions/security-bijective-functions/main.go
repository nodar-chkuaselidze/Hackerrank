package main

import "fmt"

func main() {
	var n int
	var set map[int]int

	fmt.Scan(&n)
	set = make(map[int]int)

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		set[x]++
	}

	for _, v := range set {
		if v > 1 {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
