package main

import (
	"fmt"
	"math"
)

func main() {
	var T uint
	var N uint

	fmt.Scan(&T)

	for i := uint(0); i < T; i++ {
		fmt.Scan(&N)
		fmt.Println(calculateUtopianTree(N))
	}
}

func calculateUtopianTree(cycles uint) uint {
	if cycles == 0 {
		return 1
	}

	power := 2 + ((cycles - 1) / 2)

	return uint(math.Pow(2, float64(power))) - (1 + (cycles % 2))
}
