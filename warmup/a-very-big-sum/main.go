package main

import (
	"fmt"
)

func main() {
	var count int

	var sum, tmp int64

	sum = 0

	fmt.Scan(&count)

	for i := 0; i < count; i++ {
		fmt.Scan(&tmp)
		sum += tmp
	}

	fmt.Println(sum)
}
