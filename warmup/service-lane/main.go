package main

import "fmt"

func main() {
	var N, T uint
	var Widths []uint

	fmt.Scanln(&N, &T)

	Widths = make([]uint, N)
	for i := uint(0); i < N; i++ {
		fmt.Scan(&Widths[i])
	}

	var startIdx, endIdx uint
	for i := uint(0); i < T; i++ {
		fmt.Scanln(&startIdx, &endIdx)
		fmt.Println(min(Widths, startIdx, endIdx))
	}
}

func min(Widths []uint, i, j uint) uint {
	var min uint

	min = Widths[i]
	for k := i + 1; k <= j; k++ {
		if min > Widths[k] {
			min = Widths[k]
		}
	}

	return min
}
