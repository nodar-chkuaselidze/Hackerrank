package main

import "fmt"

func main() {
	var S string

	fmt.Scanln(&S)

	if canBePalindrome(S) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func canBePalindrome(S string) bool {
	Counts := make(map[uint8]int)

	for i, length := 0, len(S); i < length; i++ {
		if _, ok := Counts[S[i]]; !ok {
			Counts[S[i]] = 0
		}

		Counts[S[i]]++
	}

	foundOdd := false
	for _, count := range Counts {
		isOdd := count%2 == 1

		if isOdd && !foundOdd {
			foundOdd = true
		} else if isOdd {
			return false
		}
	}

	return true
}
