package main

import (
	"fmt"
	"math"
)

func main() {
	var matrixSize int
	sumLeft := 0
	sumRight := 0

	fmt.Scan(&matrixSize)

	matrix := make([][]int, matrixSize)

	for i := 0; i < matrixSize; i++ {
		matrix[i] = make([]int, matrixSize)

		for j := 0; j < matrixSize; j++ {
			fmt.Scan(&matrix[i][j])
			if i == j {
				sumLeft += matrix[i][j]
			}

			if matrixSize-j-1 == i {
				sumRight += matrix[i][j]
			}
		}
	}

	fmt.Println(math.Abs(float64(sumLeft) - float64(sumRight)))

}
