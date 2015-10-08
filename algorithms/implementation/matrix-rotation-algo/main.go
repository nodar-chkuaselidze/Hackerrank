package main

import (
	"fmt"
)

func main() {
	var M, N, R int

	fmt.Scan(&M, &N, &R)

	matrix := make([][]int, M)

	for i := 0; i < M; i++ {
		matrix[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	//printMatrix(matrix)

	layersCount := min(M, N) / 2

	xmin := 0
	ymin := 0

	for i := 0; i < layersCount; i++ {
		layerLength := ((N - xmin) * 2) + ((M - ymin) * 2) - 4

		for j := 0; j < R%layerLength; j++ {
			rotateLayerOnce(matrix, xmin, ymin, M, N)
		}
		M--
		N--
		xmin++
		ymin++
	}

	//fmt.Println("----")
	printMatrix(matrix)
}

func rotateLayerOnce(matrix [][]int, xmin, ymin, A, B int) {
	xMin := xmin
	xMax := B
	yMin := ymin
	yMax := A

	tmp := matrix[xMin][yMin]

	//first line
	for i := xMin; i < xMax-1; i++ {
		matrix[yMin][i] = matrix[yMin][i+1]
	}

	//right side
	for i := yMin; i < yMax-1; i++ {
		matrix[i][xMax-1] = matrix[i+1][xMax-1]
	}

	//bottom side
	for i := xMax - 1; i > xMin; i-- {
		matrix[yMax-1][i] = matrix[yMax-1][i-1]
	}

	//left side
	for i := yMax - 1; i > yMin+1; i-- {
		matrix[i][xMin] = matrix[i-1][xMin]
	}

	matrix[yMin+1][xMin] = tmp
}

func printMatrix(matrix [][]int) {
	xLen := len(matrix)

	for i := 0; i < xLen; i++ {
		yLen := len(matrix[i])
		for j := 0; j < yLen; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println("")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
