package main

import "fmt"

func main() {
	var n int
	var inv []int

	fmt.Scan(&n)
	inv = make([]int, n)

	for i := range inv {
		fmt.Scan(&inv[i])
	}

	for i, v := range inv {
		if inv[v-1] != i+1 {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
