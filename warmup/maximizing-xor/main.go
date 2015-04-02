package main

import (
	"fmt"
	"math"
)

func main() {
	var L, R, xor, i uint

	fmt.Scan(&L, &R)

	xor = L ^ R

	i = 0

	for xor > uint(math.Pow(2, float64(i))) {
		i++
	}

	fmt.Print(math.Pow(2, float64(i)) - 1)
}
