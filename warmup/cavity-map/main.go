package main

import "fmt"

func main() {
	var N int
	var lines [3][]byte

	fmt.Scan(&N)

	for i := 0; i < 3; i++ {
		lines[i] = make([]byte, N)
	}

	fmt.Scanln(&lines[0])
	fmt.Println(string(lines[0]))

	if N < 2 {
		return
	}

	fmt.Scanln(&lines[1])

	if N == 2 {
		fmt.Println(string(lines[1]))
		return
	}

	for i := 2; i < N; i++ {
		fmt.Scanln(&lines[2])

		fmt.Println(replace(lines))

		lines[0] = lines[1]
		lines[1] = lines[2]
	}

	fmt.Println(string(lines[2]))
}

func replace(lines [3][]byte) string {
	length := len(lines[1])
	printLine := make([]byte, length)

	printLine[0] = lines[1][0]
	printLine[length-1] = lines[1][length-1]

	for i := 1; i < len(lines[1])-1; i++ {
		current := lines[1][i]

		if current > lines[0][i] && current > lines[2][i] && current > lines[1][i-1] && current > lines[1][i+1] {
			printLine[i] = 'X'
		} else {
			printLine[i] = current
		}
	}

	return string(printLine)
}
