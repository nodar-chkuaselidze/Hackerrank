package main

import (
	"fmt"
	"math"
)

func main() {
	var T int
	var S string

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scanln(&S)
		palindrome := palindromize([]byte(S))

		fmt.Println(palindrome)
	}
}

func palindromize(S []byte) uint {
	count := 0
	length := len(S)
	lenFloat := float64(length)

	for i := int(math.Ceil(lenFloat / 2)); i < length; i++ {
		count += int(math.Abs(float64(S[i]) - float64(S[length-i-1])))
	}

	return uint(count)
}
