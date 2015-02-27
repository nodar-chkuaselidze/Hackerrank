package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, K int

	fmt.Scan(&N, &K)
	readNumbers := make([]int, 0, N)

	for i := 0; i < N; i++ {
		var n int

		fmt.Scan(&n)
		readNumbers = append(readNumbers, n)
	}

	numbers := sort.IntSlice(readNumbers)
	numbers.Sort()

	var unfairness int

	unfairness = maxMin(numbers[0:K])

	for i := 1; i <= N-K; i++ {
		unfair := maxMin(numbers[i : K+i])

		if unfairness > unfair {
			unfairness = unfair
		}
	}

	fmt.Println(unfairness)
}

func maxMin(X []int) int {
	return X[len(X)-1] - X[0]
}
