package main

import (
	"fmt"
	"math"
)

func main() {
	var N, M, sum uint

	fmt.Scan(&N, &M)

	sum = 0
	for i := uint(0); i < M; i++ {
		var a, b, k uint
		fmt.Scanln(&a, &b, &k)

		sum += (b - a + 1) * k
	}

	fmt.Println(uint(math.Floor(float64(sum / N))))
}
