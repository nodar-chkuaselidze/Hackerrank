package main

import "fmt"
import "math"

func main() {
	var text string

	fmt.Scanf("%s", &text)

	flen := float64(len(text))
	sqrt := math.Sqrt(flen)
	low := math.Floor(sqrt)
	high := math.Ceil(sqrt)

	if low*high < flen {
		low = high
	}

	fmt.Println(encrypt(text, int(low), int(high)))
}

func encrypt(text string, rows, columns int) (ctext string) {
	l := len(text)
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			idx := i + (columns * j)

			if idx >= l {
				continue
			}

			ctext += string(text[idx])
		}

		ctext += " "
	}
	return
}
