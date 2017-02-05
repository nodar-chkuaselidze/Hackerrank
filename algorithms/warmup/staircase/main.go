package main

import "fmt"
import "strings"

const PAD = " "
const STAIR = "#"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	for i := 1; i < n+1; i++ {
		fmt.Println(strings.Repeat(PAD, n-i) + strings.Repeat(STAIR, i))
	}
}
