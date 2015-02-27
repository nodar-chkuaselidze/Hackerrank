package main

import "fmt"

func main() {
	var T uint
	var S string

	fmt.Scan(&T)

	for i := T; i > 0; i-- {
		fmt.Scanln(&S)
		fmt.Println(deletions(S))
	}
}

func deletions(S string) uint {
	count := uint(0)
	latest := S[0]

	for i := 1; i < len(S); i++ {
		if latest == S[i] {
			count++
		}

		latest = S[i]
	}

	return count
}
