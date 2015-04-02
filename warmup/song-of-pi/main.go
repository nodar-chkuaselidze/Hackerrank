package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var PI_Numbers []int

func main() {
	var T uint
	var line string

	PI_Numbers = func(numbers []string) []int {
		nums := make([]int, len(numbers))

		for i, val := range numbers {
			nums[i], _ = strconv.Atoi(val)
		}

		return nums
	}(strings.Split("31415926535897932384626433833", ""))

	fmt.Scanln(&T)

	scanner := bufio.NewScanner(os.Stdin)

	i := uint(0)
	for scanner.Scan() && i < T {
		i++
		line = scanner.Text()
		fmt.Println("It's " + isPIstr(line) + "a pi song.")
	}
}

func isPIstr(str string) string {
	words := strings.Split(str, " ")

	for i := 0; i < len(words); i++ {
		if len(words[i]) != PI_Numbers[i] {
			return "not "
		}
	}

	return ""
}
