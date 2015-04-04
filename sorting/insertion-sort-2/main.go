package main

import "fmt"

func main() {
	var N int

	fmt.Scanln(&N)

	array := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&array[i])
	}

	if N > 1 {
		sortArray(array, 2)
	}
}

func sortArray(array []int, index int) {
	N := index
	unsorted := array[N-1]

	for i := N - 2; i >= -1; i-- {
		if i > -1 && array[i] > unsorted {
			array[i+1] = array[i]
		} else {
			array[i+1] = unsorted
			break
		}
	}
	printArray(array)
	if index < len(array) {
		sortArray(array, index+1)
	}
}

func printArray(array []int) {
	for i := 0; i < len(array); i++ {
		fmt.Print(array[i], " ")
	}
	fmt.Print("\n")
}
