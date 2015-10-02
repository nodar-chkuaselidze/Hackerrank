package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int

	fmt.Scanln(&N)

	sticks := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&sticks[i])
	}

	sort.Ints(sticks)

	sticksCount := cutSticks(sticks)

	fmt.Println(N)
	for sticksCount > 0 {
		fmt.Println(sticksCount)
		sticksCount = cutSticks(sticks)
	}
}

func cutSticks(sticks []int) int {
	var count, min int = 0, 0

	for i, value := range sticks {
		if value == 0 {
			continue
		}

		if min == 0 {
			min = value
		}

		sticks[i] -= min
		if sticks[i] > 0 {
			count++
		}
	}

	return count
}
