package main

import (
	"fmt"
	"math"
)

func main() {
	var T, N, C, M int

	fmt.Scan(&T)
	for res, i := 0.0, 0; i < T; i++ {
		fmt.Scanln(&N, &C, &M)

		res = math.Floor(float64(N / C))

		res += freeChocolates(float64(M), res)
		fmt.Println(uint(res))
	}
}

func freeChocolates(M, res float64) float64 {
	if M > res {
		return 0
	}

	chocolates := math.Floor(res / M)
	return chocolates + freeChocolates(M, (res-chocolates*M)+chocolates)
}
