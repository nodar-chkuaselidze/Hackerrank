package main

import "fmt"

func main() {
	var count, sum, tmp int

	fmt.Scan(&count)

	sum = 0
	for i := 0; i < count; i++ {
		fmt.Scan(&tmp)
		sum += tmp
	}

	fmt.Println(sum)
}
