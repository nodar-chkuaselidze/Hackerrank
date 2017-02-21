package main

import "fmt"

func main() {
	var year int

	fmt.Scan(&year)

	date := 13

	if year == 1918 {
		date = 27
	}

	if year > 1918 {
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			date = 12
		}
	}

	if year < 1918 && year%4 == 0 {
		date = 12
	}

	fmt.Printf("%d.09.%d", date, year)
}
