package main

import (
	"fmt"
)

func main() {
	var count float32

	var countZeros, countNegative, countPositive float32

	fmt.Scan(&count)

	for i := float32(0); i < count; i++ {
		var tmp float32
		fmt.Scan(&tmp)

		if tmp == 0 {
			countZeros++
			continue
		}

		if tmp < 0 {
			countNegative++
		} else {
			countPositive++
		}
	}

	fmt.Println(trunc(countPositive / count))
	fmt.Println(trunc(countNegative / count))
	fmt.Println(trunc(countZeros / count))
}

func trunc(number float32) float32 {
	return float32(int(number*1000)) / 1000
}
