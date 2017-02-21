package main

import "fmt"
import "sort"

type bigNums []string

func (b bigNums) Len() int      { return len(b) }
func (b bigNums) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func (b bigNums) Less(i, j int) bool {
	if len(b[i]) == len(b[j]) {
		bi := b[i]
		bj := b[j]

		for i := range bi {
			if bi[i] == bj[i] {
				continue
			}

			return bi[i] < bj[i]
		}
	}

	return len(b[i]) < len(b[j])
}

func main() {
	var n int
	var unsorted bigNums

	fmt.Scan(&n)

	unsorted = make(bigNums, n)

	for i := range unsorted {
		fmt.Scanln(&unsorted[i])
	}

	sort.Sort(unsorted)

	for _, v := range unsorted {
		fmt.Println(v)
	}
}
