package main

import "fmt"

func main() {
	var heights [26]int
	var text string

	for i := range heights {
		fmt.Scan(&heights[i])
	}

	fmt.Scan(&text)

	max := 1
	for _, char := range text {
		i := char - 97
		if max <= heights[i] {
			max = heights[i]
		}
	}

	fmt.Println(len(text) * max)
}
