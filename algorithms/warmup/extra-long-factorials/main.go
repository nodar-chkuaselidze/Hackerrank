package main

import (
	"fmt"
	"math/big"
)

func main() {
	var factSize int64

	fmt.Scan(&factSize)

	fmt.Println(factorial(factSize))
}

func factorial(number int64) *big.Int {
	if number == 1 {
		return big.NewInt(1)
	}

	temp := new(big.Int)
	prevFactorial := factorial(number - 1)
	return temp.Mul(prevFactorial, big.NewInt(number))
}
