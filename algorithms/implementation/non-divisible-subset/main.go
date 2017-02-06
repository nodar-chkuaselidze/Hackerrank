package main

import "fmt"

func main() {
	var n, k int
	var remainders []int

	fmt.Scan(&n, &k)

	remainders = make([]int, k)

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		remainders[x%k]++
	}

	var max int
	if remainders[0] > 0 {
		max = 1
	} else {
		max = 0
	}

	for i := 1; i < k/2+1; i++ {
		if remainders[i] > remainders[k-i] {
			max += remainders[i]
		} else {
			max += remainders[k-i]
		}
	}

	if k%2 == 0 && remainders[k/2] > 1 {
		max -= remainders[k/2] - 1
	}

	fmt.Println(max)
}
